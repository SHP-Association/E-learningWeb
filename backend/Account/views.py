# myapp/views.py (or accounts/views.py)

from django.contrib.auth import get_user_model
from django.contrib.auth.tokens import default_token_generator
from django.utils.encoding import force_str, force_bytes
from django.utils.http import urlsafe_base64_decode, urlsafe_base64_encode
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt # For API, consider DRF or more robust CSRF handling
from django.conf import settings
from django.template.loader import render_to_string
from django.core.exceptions import ValidationError
from django.contrib.auth.password_validation import validate_password
from django.contrib.auth import update_session_auth_hash # To keep user logged in after password change (optional)
# from Account.views import password_reset_request_api, password_reset_confirm_api
import json
import os
# Import your custom mail utility
# Ensure 'myapp.mail' matches the actual path to your mail.py file
from utils.mail import trigger_email
from django.core.mail import send_mail, BadHeaderError

User = get_user_model()

@csrf_exempt # IMPORTANT: For production, use Django REST Framework or proper CSRF token handling.
def password_reset_request_api(request):
    """
    API endpoint to initiate the password reset process.
    Expects a POST request with 'email' in the JSON body.
    Sends a password reset link to the user's email.
    """
    if request.method == 'POST':
        try:
            data = json.loads(request.body)
            email = data.get('email')
        except json.JSONDecodeError:
            return JsonResponse({'error': 'Invalid JSON body.'}, status=400)

        if not email:
            return JsonResponse({'error': 'Email field is required.'}, status=400)

        try:
            user = User.objects.get(email=email)
        except User.DoesNotExist:
            # For security, always return a generic success message
            # to prevent revealing whether an email address exists in the system.
            return JsonResponse({'message': 'If an account with that email exists, a password reset email has been sent.'}, status=200)

        # Use settings.FRONTEND_URL for reliability
        # frontend_url = getattr(settings, 'FRONTEND_URL', None)
        frontend_url = os.getenv("FRONTEND_URL")
        if not frontend_url:
            print("FRONTEND_URL is not set in Django settings.")
            return JsonResponse({'error': 'Server misconfiguration: FRONTEND_URL not set.'}, status=500)

        # Generate unique ID (uidb64) and token for the reset link
        uid = urlsafe_base64_encode(force_bytes(user.pk))
        token = default_token_generator.make_token(user)

        # Construct the full password reset URL for the frontend
        # This URL MUST point to your React frontend's PasswordResetConfirm page.
        # Example: http://localhost:5173/password_reset/<uidb64>/<token>/
        # Ensure settings.FRONTEND_URL is correctly configured in your Django settings.
        frontend_reset_url = f"{frontend_url.rstrip('/')}/password_reset/{uid}/{token}/"

        # Prepare context for the email template
        email_context = {
            'user': user,
            'protocol': os.getenv('EMAIL_PROTOCOL', 'http'),  # Use env or default to 'http'
            'domain': frontend_url,  # Your frontend domain for email link display
            'uid': uid,
            'token': token,
            'reset_url': frontend_reset_url,  # The full URL to be used in the email template
        }

        # Debug print for troubleshooting
        print("Sending password reset email to:", user.email)
        print("Email context:", email_context)

        try:
            # Call your custom email utility function
            email_sent_error = trigger_email(
                context=email_context,
                template='emails/password_reset_email.html', # Fixed template path
                subject='Password Reset for your SHP-Learner account',
                recipients=[user.email],
                message="Please use the link below to reset your password." # Plain text fallback
            )
        except Exception as e:
            print("Exception in trigger_email:", e)
            return JsonResponse({'error': f'Failed to send password reset email: {str(e)}'}, status=500)

        if email_sent_error:
            # Log the error for debugging, and return a more explicit error for troubleshooting
            print(f"Error sending password reset email to {user.email}: {email_sent_error}")
            return JsonResponse({'error': f'Failed to send password reset email: {email_sent_error}'}, status=500)

        return JsonResponse({'message': 'If an account with that email exists, a password reset email has been sent.'}, status=200)
    else:
        return JsonResponse({'error': 'Only POST requests are allowed for this endpoint.'}, status=405)


@csrf_exempt # IMPORTANT: For production, use Django REST Framework or proper CSRF token handling.
def password_reset_confirm_api(request, uidb64, token):
    """
    API endpoint to confirm password reset and set the new password.
    Expects a POST request with 'new_password1' and 'new_password2' in the JSON body,
    and uidb64 and token from the URL path.
    """
    if request.method == 'POST':
        try:
            data = json.loads(request.body)
            new_password1 = data.get('new_password1')
            new_password2 = data.get('new_password2')
        except json.JSONDecodeError:
            return JsonResponse({'error': 'Invalid JSON body.'}, status=400)

        if not all([new_password1, new_password2]):
            return JsonResponse({'error': 'Both new password fields are required.'}, status=400)

        if new_password1 != new_password2:
            return JsonResponse({'error': 'Passwords do not match.'}, status=400)

        user = None
        try:
            # Decode uidb64 to get the user ID
            uid = force_str(urlsafe_base64_decode(uidb64))
            user = User.objects.get(pk=uid)
        except (TypeError, ValueError, OverflowError, User.DoesNotExist):
            user = None # User not found or uid is invalid

        # Check if user exists and the token is valid
        if user is not None and default_token_generator.check_token(user, token):
            # Validate password complexity using Django's built-in validators
            try:
                validate_password(new_password1, user=user)
            except ValidationError as e:
                # Return validation errors to the frontend
                return JsonResponse({'error': list(e.messages)}, status=400)

            # Set the new password and save the user
            user.set_password(new_password1)
            user.save()

            # Optional: Update the user's session hash to keep them logged in
            # if they were authenticated when they requested the reset.
            # If you want to force a logout, remove this line.
            update_session_auth_hash(request, user)

            return JsonResponse({'message': 'Password has been reset successfully.'}, status=200)
        else:
            # If user is None or token is invalid/expired
            return JsonResponse({'error': 'The password reset link is invalid or has expired.'}, status=400)
    else:
        return JsonResponse({'error': 'Only POST requests are allowed for this endpoint.'}, status=405)

# NOTE:
# Make sure your 'registration/password_reset_email.html' template uses {{ reset_url }} for the reset link.
# Example:
# <a href="{{ reset_url }}">Reset your password</a>
# Do NOT use {% url 'password_reset_confirm' uidb64=uid token=token %} in the template,
# because that only works with Django's built-in views and URL names.
# because that only works with Django's built-in views and URL names.
