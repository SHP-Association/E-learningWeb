from django.test import TestCase
from .models import Course

class CourseModelTest(TestCase):

    def setUp(self):
        Course.objects.create(title="Test Course", description="Test Description")

    def test_course_creation(self):
        course = Course.objects.get(title="Test Course")
        self.assertEqual(course.description, "Test Description")
        self.assertEqual(course.title, "Test Course")