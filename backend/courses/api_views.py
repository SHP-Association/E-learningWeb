from rest_framework import generics
from rest_framework.permissions import IsAuthenticated, AllowAny, IsAdminUser
from rest_framework.authentication import SessionAuthentication, BasicAuthentication
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from django.contrib.auth import authenticate, login
from django.contrib.auth.password_validation import validate_password
from rest_framework.exceptions import ValidationError

from .models import Course
from Enrollment.models import Enrollment
from Account.models import CustomUser
from Category.models import Category
from Lesson.models import Lesson
from Quiz.models import Quiz
from FAQ.models import FAQ
from .serializers import (
    CustomUserSerializer, CategorySerializer, CourseSerializer,
    LessonSerializer, EnrollmentSerializer, QuizSerializer, FAQSerializer
)

import os
import logging
import traceback
from utils.mail import trigger_email

logger = logging.getLogger(__name__)


# --- Custom Permission Classes ---
class IsOwnerOrAdmin(object):
    """
    Custom permission to only allow owners of an object or admins to edit it.
    """
    def has_object_permission(self, request, view, obj):
        # Admin users have full access
        if request.user.is_staff or getattr(request.user, 'role', None) == 'admin':
            return True
        # Object owner has access
        return obj == request.user


# --- Custom Permission Helpers ---
def is_instructor(user):
    return user.is_authenticated and getattr(user, "role", None) == "instructor"


def is_admin(user):
    return user.is_authenticated and (getattr(user, "role", None) == "admin" or user.is_superuser)


def is_student(user):
    return user.is_authenticated and getattr(user, "role", None) == "student"


# --- CustomUser API Views ---
class CustomUserListCreateAPIView(generics.ListCreateAPIView):
    """
    API view to list all CustomUsers or create a new CustomUser.
    - POST (create): Anyone can register
    - GET (list): Only admins can list all users
    """
    serializer_class = CustomUserSerializer
    authentication_classes = [SessionAuthentication, BasicAuthentication]

    def get_queryset(self):
        # Only admins can list all users
        if self.request.user.is_staff:
            return CustomUser.objects.all().order_by('-date_joined')
        # Non-admins can only see themselves
        if self.request.user.is_authenticated:
            return CustomUser.objects.filter(id=self.request.user.id).order_by('-date_joined')
        return CustomUser.objects.none()

    def get_permissions(self):
        if self.request.method == 'POST':  # Registration
            return [AllowAny()]
        return [IsAuthenticated()]  # List requires authentication


class CustomUserRetrieveUpdateDestroyAPIView(generics.RetrieveUpdateDestroyAPIView):
    """
    API view to retrieve, update, or delete a specific CustomUser by ID.
    Users can only access their own profile unless they're admin.
    """
    queryset = CustomUser.objects.all()
    serializer_class = CustomUserSerializer
    permission_classes = [IsAuthenticated]
    authentication_classes = [SessionAuthentication, BasicAuthentication]

    def get_object(self):
        obj = super().get_object()
        # Users can only access their own profile unless admin
        if obj != self.request.user and not is_admin(self.request.user):
            from rest_framework.exceptions import PermissionDenied
            raise PermissionDenied("You can only access your own profile.")
        return obj

    def patch(self, request, *args, **kwargs):
        user = self.get_object()
        if request.user != user and not is_admin(request.user):
            return Response({'detail': 'Not allowed.'}, status=status.HTTP_403_FORBIDDEN)
        return super().partial_update(request, *args, **kwargs)

    def delete(self, request, *args, **kwargs):
        user = self.get_object()
        # Only admin can delete users
        if not is_admin(request.user):
            return Response({'detail': 'Not allowed.'}, status=status.HTTP_403_FORBIDDEN)
        return super().delete(request, *args, **kwargs)


# --- Category API Views ---
class CategoryListCreateAPIView(generics.ListCreateAPIView):
    """
    API view to list all Categories or create a new Category.
    - GET: Anyone can view
    - POST: Only instructors/admins can create
    """
    queryset = Category.objects.all().order_by('name')
    serializer_class = CategorySerializer

    def get_permissions(self):
        if self.request.method == 'POST':
            return [IsAuthenticated()]
        return [AllowAny()]

    def perform_create(self, serializer):
        if not (is_admin(self.request.user) or is_instructor(self.request.user)):
            from rest_framework.exceptions import PermissionDenied
            raise PermissionDenied("Only instructors and admins can create categories.")
        serializer.save()


class CategoryRetrieveUpdateDestroyAPIView(generics.RetrieveUpdateDestroyAPIView):
    """
    API view to retrieve, update, or delete a specific Category by ID.
    """
    queryset = Category.objects.all()
    serializer_class = CategorySerializer

    def get_permissions(self):
        if self.request.method == 'GET':
            return [AllowAny()]
        return [IsAuthenticated()]

    def put(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().put(request, *args, **kwargs)

    def patch(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().patch(request, *args, **kwargs)

    def delete(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().delete(request, *args, **kwargs)


# --- Course API Views ---
class CourseListCreateAPIView(generics.ListCreateAPIView):
    """
    API view to list all Courses or create a new Course.
    - GET: Anyone can view published courses
    - POST: Only instructors/admins can create
    """
    serializer_class = CourseSerializer

    def get_queryset(self):
        # Optimize with select_related and prefetch_related to avoid N+1 queries
        queryset = Course.objects.select_related(
            'instructor', 'category'
        ).prefetch_related('lessons').order_by('-created_at')
        
        # Non-staff users only see published courses
        if not self.request.user.is_staff:
            queryset = queryset.filter(is_published=True)
        
        return queryset

    def get_permissions(self):
        if self.request.method == 'POST':
            return [IsAuthenticated()]
        return [AllowAny()]

    def perform_create(self, serializer):
        if not (is_admin(self.request.user) or is_instructor(self.request.user)):
            from rest_framework.exceptions import PermissionDenied
            raise PermissionDenied("Only instructors and admins can create courses.")
        serializer.save(instructor=self.request.user)


class CourseRetrieveUpdateDestroyAPIView(generics.RetrieveUpdateDestroyAPIView):
    """
    API view to retrieve, update, or delete a specific Course by ID.
    """
    serializer_class = CourseSerializer

    def get_queryset(self):
        queryset = Course.objects.select_related(
            'instructor', 'category'
        ).prefetch_related('lessons')
        
        # Non-staff users only see published courses
        if not self.request.user.is_staff:
            queryset = queryset.filter(is_published=True)
        
        return queryset

    def get_permissions(self):
        if self.request.method == 'GET':
            return [AllowAny()]
        return [IsAuthenticated()]

    def put(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().put(request, *args, **kwargs)

    def patch(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().patch(request, *args, **kwargs)

    def delete(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().delete(request, *args, **kwargs)


# --- Lesson API Views ---
class LessonListCreateAPIView(generics.ListCreateAPIView):
    """
    API view to list all Lessons or create a new Lesson.
    """
    queryset = Lesson.objects.all().order_by('order')
    serializer_class = LessonSerializer

    def get_permissions(self):
        if self.request.method == 'POST':
            return [IsAuthenticated()]
        return [AllowAny()]

    def perform_create(self, serializer):
        if not (is_admin(self.request.user) or is_instructor(self.request.user)):
            from rest_framework.exceptions import PermissionDenied
            raise PermissionDenied("Only instructors and admins can create lessons.")
        serializer.save()


class LessonRetrieveUpdateDestroyAPIView(generics.RetrieveUpdateDestroyAPIView):
    """
    API view to retrieve, update, or delete a specific Lesson by ID.
    """
    queryset = Lesson.objects.all()
    serializer_class = LessonSerializer

    def get_permissions(self):
        if self.request.method == 'GET':
            return [AllowAny()]
        return [IsAuthenticated()]

    def put(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().put(request, *args, **kwargs)

    def patch(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().patch(request, *args, **kwargs)

    def delete(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().delete(request, *args, **kwargs)


# --- Enrollment API Views ---
class EnrollmentListCreateAPIView(generics.ListCreateAPIView):
    """
    API view to list all Enrollments or create a new Enrollment.
    Users can only see their own enrollments unless they're admin.
    """
    serializer_class = EnrollmentSerializer
    authentication_classes = [SessionAuthentication, BasicAuthentication]
    permission_classes = [IsAuthenticated]

    def get_queryset(self):
        user = self.request.user
        logger.info(f"User {user.username} is requesting enrollments")
        
        # Admins can see all enrollments
        if user.is_staff:
            return Enrollment.objects.all().order_by('-enrolled_at')
        
        # Regular users only see their own enrollments
        return Enrollment.objects.filter(student=user).order_by('-enrolled_at')

    def perform_create(self, serializer):
        serializer.save(student=self.request.user)


class EnrollmentRetrieveUpdateDestroyAPIView(generics.RetrieveUpdateDestroyAPIView):
    """
    API view to retrieve, update, or delete a specific Enrollment by ID.
    Only authenticated users (typically the student or an admin) can access.
    """
    queryset = Enrollment.objects.all()
    serializer_class = EnrollmentSerializer
    permission_classes = [IsAuthenticated]
    authentication_classes = [SessionAuthentication, BasicAuthentication]

    def get_object(self):
        obj = super().get_object()
        # Users can only access their own enrollments unless admin
        if obj.student != self.request.user and not is_admin(self.request.user):
            from rest_framework.exceptions import PermissionDenied
            raise PermissionDenied("You can only access your own enrollments.")
        return obj

    def delete(self, request, *args, **kwargs):
        # Only admin can delete enrollments
        if not is_admin(request.user):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().delete(request, *args, **kwargs)


# --- Quiz API Views ---
class QuizListCreateAPIView(generics.ListCreateAPIView):
    """
    API view to list all Quizzes or create a new Quiz.
    """
    queryset = Quiz.objects.all().order_by('-created_at')
    serializer_class = QuizSerializer

    def get_permissions(self):
        if self.request.method == 'POST':
            return [IsAuthenticated()]
        return [AllowAny()]

    def perform_create(self, serializer):
        if not (is_admin(self.request.user) or is_instructor(self.request.user)):
            from rest_framework.exceptions import PermissionDenied
            raise PermissionDenied("Only instructors and admins can create quizzes.")
        serializer.save()


class QuizRetrieveUpdateDestroyAPIView(generics.RetrieveUpdateDestroyAPIView):
    """
    API view to retrieve, update, or delete a specific Quiz by ID.
    """
    queryset = Quiz.objects.all()
    serializer_class = QuizSerializer

    def get_permissions(self):
        if self.request.method == 'GET':
            return [AllowAny()]
        return [IsAuthenticated()]

    def put(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().put(request, *args, **kwargs)

    def patch(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().patch(request, *args, **kwargs)

    def delete(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().delete(request, *args, **kwargs)


# --- FAQ API Views ---
class FAQListCreateAPIView(generics.ListCreateAPIView):
    """
    API view to list all FAQs or create a new FAQ.
    Only published FAQs are shown to non-staff users.
    """
    serializer_class = FAQSerializer

    def get_queryset(self):
        if self.request.user.is_staff:
            return FAQ.objects.all().order_by('-created_at')
        return FAQ.objects.filter(is_published=True).order_by('-created_at')

    def get_permissions(self):
        if self.request.method == 'POST':
            return [AllowAny()]  # Anyone can submit questions
        return [AllowAny()]


class FAQRetrieveUpdateDestroyAPIView(generics.RetrieveUpdateDestroyAPIView):
    """
    API view to retrieve, update, or delete a specific FAQ by ID.
    """
    queryset = FAQ.objects.all()
    serializer_class = FAQSerializer

    def get_permissions(self):
        if self.request.method == 'GET':
            return [AllowAny()]
        return [IsAuthenticated()]

    def put(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().put(request, *args, **kwargs)

    def patch(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().patch(request, *args, **kwargs)

    def delete(self, request, *args, **kwargs):
        if not (is_admin(request.user) or is_instructor(request.user)):
            return Response({'detail': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)
        return super().delete(request, *args, **kwargs)


# --- Login API View ---
class LoginAPIView(APIView):
    """
    API endpoint for logging in a user using username and password.
    """
    permission_classes = [AllowAny]

    def post(self, request, *args, **kwargs):
        username = request.data.get('username')
        password = request.data.get('password')

        if not username or not password:
            return Response(
                {'detail': 'Username and password are required'}, 
                status=status.HTTP_400_BAD_REQUEST
            )

        user = authenticate(request, username=username, password=password)

        if user is not None and user.is_active:
            login(request, user)
            logger.info(f"User {username} logged in successfully")
            return Response({'detail': 'Login successful'}, status=status.HTTP_200_OK)

        logger.warning(f"Failed login attempt for username: {username}")
        return Response({'detail': 'Invalid credentials'}, status=status.HTTP_401_UNAUTHORIZED)


# --- Register API View ---
class RegisterAPIView(APIView):
    """
    API endpoint for registering a new user.
    """
    permission_classes = [AllowAny]

    def post(self, request, *args, **kwargs):
        username = request.data.get('username')
        email = request.data.get('email')
        password = request.data.get('password')
        role = request.data.get('role', 'student')

        # Basic validation
        if not username or not email or not password:
            return Response({'message': 'All fields are required.'}, status=status.HTTP_400_BAD_REQUEST)

        if CustomUser.objects.filter(username=username).exists():
            return Response({'message': 'Username already taken.'}, status=status.HTTP_400_BAD_REQUEST)

        if CustomUser.objects.filter(email=email).exists():
            return Response({'message': 'Email already registered.'}, status=status.HTTP_400_BAD_REQUEST)

        try:
            validate_password(password)
        except Exception as e:
            return Response({'message': str(e)}, status=status.HTTP_400_BAD_REQUEST)

        # Create user with role
        user = CustomUser.objects.create_user(
            username=username,
            email=email,
            password=password,
            role=role
        )
        # The save method will auto-assign is_staff based on role
        user.save()

        logger.info(f"New user registered: {username} with role: {role}")

        # --- Send Welcome Email ---
        try:
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
                subject=f'Welcome to SHP-Learner, {user.username}!',
                recipients=[user.email],
                message=f"Welcome to SHP-Learner, {user.username}! We're excited to have you."
            )

            if email_error:
                logger.error(f"Failed to send welcome email to {user.email}: {email_error}")

        except Exception as e:
            logger.exception(f"Unexpected error when trying to send welcome email to {user.email}")

        return Response({'message': 'Registration successful!'}, status=status.HTTP_201_CREATED)