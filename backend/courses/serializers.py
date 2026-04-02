from rest_framework import serializers
from .models import Course
from Enrollment.models import Enrollment
from Account.models import CustomUser
from Category.models import Category
from Lesson.models import Lesson
from Quiz.models import Quiz
from FAQ.models import FAQ

from rest_framework import serializers
from .models import Course
from Enrollment.models import Enrollment
from Account.models import CustomUser
from Category.models import Category
from Lesson.models import Lesson
from Quiz.models import Quiz
from FAQ.models import FAQ
from Certificate.models import Certificate
from Question.models import Question, AnswerChoice
from Review.models import Review

class CustomUserSerializer(serializers.ModelSerializer):
    class Meta:
        model = CustomUser
        fields = [
            'id', 'username', 'email', 'first_name', 'last_name',
            'role', 'bio', 'profile_picture', 'date_of_birth', 'gender',
            'contact_number', 'address', 'country', 'is_email_verified',
            'highest_qualification', 'institution', 'skills',
            'linkedin_profile', 'github_profile', 'instructor_rating',
            'total_reviews', 'date_joined', 'last_login'
        ]
        read_only_fields = ['id', 'date_joined', 'last_login', 'instructor_rating', 'total_reviews']
        # Explicitly exclude sensitive fields
        # Never expose: password, is_superuser, user_permissions, groups, login_ip, last_activity

class CategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = Category
        fields = '__all__'

class LessonSerializer(serializers.ModelSerializer):
    class Meta:
        model = Lesson
        fields = '__all__'

class CourseSerializer(serializers.ModelSerializer):
    instructor = CustomUserSerializer(read_only=True)
    category = CategorySerializer(read_only=True)
    lessons = LessonSerializer(many=True, read_only=True)

    class Meta:
        model = Course
        fields = '__all__'

class EnrollmentSerializer(serializers.ModelSerializer):
    """
    Serializer for the Enrollment model.
    Includes nested serializers for student (CustomUser) and course.
    """
    student = CustomUserSerializer(read_only=True)
    course = CourseSerializer(read_only=True)

    class Meta:
        model = Enrollment
        fields = ['id', 'student', 'course', 'enrolled_at', 'progress']

class QuizSerializer(serializers.ModelSerializer):
    """
    Serializer for the Quiz model.
    Includes all fields.
    """
    class Meta:
        model = Quiz
        fields = ['id', 'lesson', 'title', 'created_at']

class FAQSerializer(serializers.ModelSerializer):
    """
    Serializer for the FAQ model.
    Includes all fields.
    """
    class Meta:
        model = FAQ
        fields = ['id', 'question', 'answer', 'created_at']

class CertificateSerializer(serializers.ModelSerializer):
    """
    Serializer for the Certificate model.
    Includes nested enrollment data.
    """
    enrollment = EnrollmentSerializer(read_only=True)
    
    class Meta:
        model = Certificate
        fields = ['id', 'enrollment', 'unique_id', 'issue_date']
        read_only_fields = ['unique_id', 'issue_date']

class AnswerChoiceSerializer(serializers.ModelSerializer):
    """
    Serializer for the AnswerChoice model.
    """
    class Meta:
        model = AnswerChoice
        fields = ['id', 'choice_text', 'is_correct']

class QuestionSerializer(serializers.ModelSerializer):
    """
    Serializer for the Question model.
    Includes nested answer choices.
    """
    choices = AnswerChoiceSerializer(many=True, read_only=True)
    
    class Meta:
        model = Question
        fields = ['id', 'quiz', 'question_text', 'question_type', 'order', 'choices']

class ReviewSerializer(serializers.ModelSerializer):
    """
    Serializer for the Review model.
    Includes nested student and course data.
    """
    student = CustomUserSerializer(read_only=True)
    course = CourseSerializer(read_only=True)
    
    class Meta:
        model = Review
        fields = ['id', 'course', 'student', 'rating', 'comment', 'created_at', 'is_approved']
        read_only_fields = ['created_at']
