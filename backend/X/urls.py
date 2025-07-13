# myproject/urls.py

from django.contrib import admin
from django.urls import path, include
from courses import views # Assuming 'courses' app views are still needed
from django.conf import settings
from django.conf.urls.static import static
# from django.contrib.auth import views as auth_views # Commented out as we are replacing with API views
from courses.views import faq_view # Assuming faq_view is still needed

# Import your new API views from your app (e.g., 'myapp')
# Make sure 'Account.views' matches the actual path to your views.py file
from Account.views import password_reset_request_api, password_reset_confirm_api

urlpatterns = [
    path('admin/', admin.site.urls),
    path('sandesh', views.index, name='index'),
    path('', views.frontend, name='index'),
    path('courses/', views.courses, name='courses'),
    path('course/<slug:slug>/', views.course_detail, name='course_detail'),
    path('profile/', views.profile, name='profile'),
    path('login/', views.user_login, name='login'),
    path('logout/', views.user_logout, name='logout'),
    path('register/', views.register, name='register'),
    path('enroll/<slug:slug>/', views.enroll, name='enroll'),
    path('api/', include('courses.urls')),
    path('faq/', faq_view, name='faq'),
    # Password Reset API Endpoints
    path('api/password_reset/request/', password_reset_request_api, name='api_password_reset_request'),
    # This endpoint matches the frontend's fetch URL: /reset/<uidb64>/<token>/
    path('reset/<str:uidb64>/<str:token>/', password_reset_confirm_api, name='api_password_reset_confirm'),
] + static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)

