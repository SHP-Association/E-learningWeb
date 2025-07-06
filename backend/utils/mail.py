# myapp/mail.py (or utils/mail.py, ensure import paths are correct)

from premailer import transform
from django.core.mail import send_mail
from django.template.loader import render_to_string
import os

def trigger_email(context, template, subject, recipients, message=None):
    """
    Sends an email using Django's send_mail, rendering an HTML template
    and inlining CSS with premailer.
    """
    try:
        # Use environment variables for all email config if available, else fallback to Django settings
        from_email = os.getenv('DEFAULT_FROM_EMAIL')
        if not from_email:
            raise Exception("DEFAULT_FROM_EMAIL is not set in environment or Django settings.")

        context['imgLogo'] = os.getenv('EMAIL_LOGO_URL')
        context['VITE_APP_BACKEND_URL'] = os.getenv('VITE_APP_BACKEND_URL')
        context['ADDRESS'] = os.getenv('ADDRESS')
        context['SUPPORT_MAIL'] = os.getenv('SUPPORT_MAIL')

        # Render the HTML template with the provided context
        raw_html = render_to_string(template_name=template, context=context)
        # Inline CSS using premailer for better email client compatibility
        html_message = transform(raw_html)

        # Send the email using Django's send_mail function
        send_mail(
            subject=subject,
            message="",  # No plain text fallback
            from_email=from_email, # Use the email configured in settings.py
            html_message=html_message, # The HTML content of the email
            recipient_list=recipients, # List of recipient email addresses
            fail_silently=False, # Set to True to suppress exceptions during email sending
        )
        return None # Return None on successful email sending
    except Exception as e:
        print(f"Error in trigger_email: {e}") # Log the error for debugging
        return e # Return the exception object on failure