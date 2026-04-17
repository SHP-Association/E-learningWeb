# X/urls.py

from django.contrib import admin
from django.urls import path, include
from courses import views
from django.conf import settings
from django.conf.urls.static import static
from courses.views import faq_view
from Account.views import PasswordResetRequestAPIView, PasswordResetConfirmAPIView

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', views.frontend, name='frontend_redirect'),
    path('courses/', views.courses, name='courses'),
    path('course/<slug:slug>/', views.course_detail, name='course_detail'),
    path('profile/', views.profile, name='profile'),
    path('login/', views.user_login, name='login'),
    path('logout/', views.user_logout, name='logout'),
    path('register/', views.register, name='register'),
    path('enroll/<slug:slug>/', views.enroll, name='enroll'),
    path('faq/', faq_view, name='faq'),
    
    # API Endpoints
    path('api/', include('courses.urls')),
    
    # Password Reset API Endpoints
    path('api/password_reset/request/', PasswordResetRequestAPIView.as_view(), name='api_password_reset_request'),
    path('reset/<str:uidb64>/<str:token>/', PasswordResetConfirmAPIView.as_view(), name='api_password_reset_confirm'),
] + static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)
