from django.urls import path
from . import views
from .views import create_course
urlpatterns = [
    path('', views.home, name='home'),
    path('courses/', views.course_list, name='course_list'),
    path('courses/<int:course_id>/', views.course_detail, name='course_detail'),
    path('courses/create/', views.create_course, name='course_create'),
]