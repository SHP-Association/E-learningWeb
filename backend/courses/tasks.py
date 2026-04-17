from celery import shared_task
from utils.mail import trigger_email
import logging

logger = logging.getLogger(__name__)

@shared_task
def send_welcome_email_task(user_id, username, email):
    """
    Background task to send a welcome email to a newly registered user.
    """
    from Account.models import CustomUser
    import os
    
    try:
        user = CustomUser.objects.get(id=user_id)
        welcome_context = {
            'user': user,
            'imgLogo': os.getenv('EMAIL_LOGO_URL'),
            'VITE_APP_BACKEND_URL': os.getenv('FRONTEND_URL'),
            'ADDRESS': os.getenv('ADDRESS'),
            'SUPPORT_MAIL': os.getenv('SUPPORT_MAIL'),
        }

        email_error = trigger_email(
            context=welcome_context,
            template='welcome_email.html',
            subject=f'Welcome to SHP-Learner, {username}!',
            recipients=[email],
            message=f"Welcome to SHP-Learner, {username}! We're excited to have you."
        )

        if email_error:
            logger.error(f"Failed to send welcome email to {email}: {email_error}")
        else:
            logger.info(f"Welcome email sent successfully to {email}")

    except CustomUser.DoesNotExist:
        logger.error(f"User with id {user_id} not found when trying to send welcome email.")
    except Exception as e:
        logger.exception(f"Unexpected error when trying to send welcome email to {email}: {e}")
