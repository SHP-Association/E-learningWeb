# Account/views.py - Password Reset API Views

from django.contrib.auth import get_user_model
from django.contrib.auth.tokens import default_token_generator
from django.utils.encoding import force_str, force_bytes
from django.utils.http import urlsafe_base64_decode, urlsafe_base64_encode
from django.core.exceptions import ValidationError
from django.contrib.auth.password_validation import validate_password
from django.contrib.auth import update_session_auth_hash
from django.core.validators import validate_email

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from rest_framework.permissions import AllowAny

import os
import logging

from utils.mail import trigger_email

logger = logging.getLogger(__name__)
User = get_user_model()


class PasswordResetRequestAPIView(APIView):
    """
    API endpoint to initiate the password reset process.
    Expects a POST request with 'email' in the JSON body.
    Sends a password reset link to the user's email.
    """
    permission_classes = [AllowAny]

    def post(self, request):
        email = request.data.get('email')

        if not email:
            return Response(
                {'error': 'Email field is required.'}, 
                status=status.HTTP_400_BAD_REQUEST
            )

        # Validate email format
        try:
            validate_email(email)
        except ValidationError:
            return Response(
                {'error': 'Invalid email format.'}, 
                status=status.HTTP_400_BAD_REQUEST
            )

        try:
            user = User.objects.get(email=email)
        except User.DoesNotExist:
            # For security, always return a generic success message
            # to prevent revealing whether an email address exists in the system.
            return Response(
                {'message': 'If an account with that email exists, a password reset email has been sent.'}, 
                status=status.HTTP_200_OK
            )

        # Get frontend URL from environment
        frontend_url = os.getenv("FRONTEND_URL")
        if not frontend_url:
            logger.error("FRONTEND_URL environment variable is not set")
            return Response(
                {'error': 'Server misconfiguration. Please contact support.'}, 
                status=status.HTTP_500_INTERNAL_SERVER_ERROR
            )

        # Generate unique ID (uidb64) and token for the reset link
        uid = urlsafe_base64_encode(force_bytes(user.pk))
        token = default_token_generator.make_token(user)

        # Construct the full password reset URL for the frontend
        frontend_reset_url = f"{frontend_url.rstrip('/')}/password_reset/{uid}/{token}/"

        # Prepare context for the email template
        email_context = {
            'user': user,
            'protocol': os.getenv('EMAIL_PROTOCOL', 'https'),
            'domain': frontend_url,
            'uid': uid,
            'token': token,
            'reset_url': frontend_reset_url,
        }

        logger.info(f"Sending password reset email to: {user.email}")

        try:
            email_sent_error = trigger_email(
                context=email_context,
                template='emails/password_reset_email.html',
                subject='Password Reset for your SHP-Learner account',
                recipients=[user.email],
                message="Please use the link below to reset your password."
            )

            if email_sent_error:
                logger.error(f"Error sending password reset email to {user.email}: {email_sent_error}")
                return Response(
                    {'error': 'Failed to send password reset email. Please try again later.'}, 
                    status=status.HTTP_500_INTERNAL_SERVER_ERROR
                )

        except Exception as e:
            logger.exception(f"Exception sending password reset email to {user.email}")
            return Response(
                {'error': 'Failed to send password reset email. Please try again later.'}, 
                status=status.HTTP_500_INTERNAL_SERVER_ERROR
            )

        return Response(
            {'message': 'If an account with that email exists, a password reset email has been sent.'}, 
            status=status.HTTP_200_OK
        )


class PasswordResetConfirmAPIView(APIView):
    """
    API endpoint to confirm password reset and set the new password.
    Expects a POST request with 'new_password1' and 'new_password2' in the JSON body,
    and uidb64 and token from the URL path.
    """
    permission_classes = [AllowAny]

    def post(self, request, uidb64, token):
        new_password1 = request.data.get('new_password1')
        new_password2 = request.data.get('new_password2')

        if not all([new_password1, new_password2]):
            return Response(
                {'error': 'Both new password fields are required.'}, 
                status=status.HTTP_400_BAD_REQUEST
            )

        if new_password1 != new_password2:
            return Response(
                {'error': 'Passwords do not match.'}, 
                status=status.HTTP_400_BAD_REQUEST
            )

        user = None
        try:
            # Decode uidb64 to get the user ID
            uid = force_str(urlsafe_base64_decode(uidb64))
            user = User.objects.get(pk=uid)
        except (TypeError, ValueError, OverflowError, User.DoesNotExist):
            user = None

        # Check if user exists and the token is valid
        if user is not None and default_token_generator.check_token(user, token):
            # Validate password complexity using Django's built-in validators
            try:
                validate_password(new_password1, user=user)
            except ValidationError as e:
                return Response(
                    {'error': list(e.messages)}, 
                    status=status.HTTP_400_BAD_REQUEST
                )

            # Set the new password and save the user
            user.set_password(new_password1)
            user.save()

            # Optional: Update the user's session hash to keep them logged in
            update_session_auth_hash(request, user)

            logger.info(f"Password reset successful for user: {user.username}")

            return Response(
                {'message': 'Password has been reset successfully.'}, 
                status=status.HTTP_200_OK
            )
        else:
            return Response(
                {'error': 'The password reset link is invalid or has expired.'}, 
                status=status.HTTP_400_BAD_REQUEST
            )


# Export the old function names for backward compatibility
password_reset_request_api = PasswordResetRequestAPIView.as_view()
password_reset_confirm_api = PasswordResetConfirmAPIView.as_view()
