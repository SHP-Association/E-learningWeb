from django.contrib.auth import authenticate, login, logout
from django.shortcuts import render, get_object_or_404, redirect
from django.contrib.auth.decorators import login_required
from django.contrib import messages
from django import forms
from django.core.mail import send_mail
from django.template.loader import render_to_string
from django.conf import settings
from .models import Course, Enrollment, CustomUser

class RegisterForm(forms.ModelForm):
    password1 = forms.CharField(widget=forms.PasswordInput, label="Password")
    password2 = forms.CharField(widget=forms.PasswordInput, label="Confirm Password")

    class Meta:
        model = CustomUser
        fields = ['username', 'email', 'role']

    def clean_password2(self):
        password1 = self.cleaned_data.get('password1')
        password2 = self.cleaned_data.get('password2')
        if password1 and password2 and password1 != password2:
            raise forms.ValidationError("Passwords don't match")
        return password2

def index(request):
    """
    Renders the homepage with a list of all courses.
    """
    courses = Course.objects.all()
    context = {'courses': courses}
    return render(request, 'index.html', context)

def courses(request):
    """
    Renders a page with all courses (similar to index for now).
    """
    courses = Course.objects.all()
    context = {'courses': courses}
    return render(request, 'index.html', context)

def course_detail(request, slug):
    """
    Renders the course detail page.
    """
    course = get_object_or_404(Course, slug=slug)
    is_enrolled = False
    if request.user.is_authenticated:
        is_enrolled = Enrollment.objects.filter(student=request.user, course=course).exists()
    context = {'course': course, 'is_enrolled': is_enrolled}
    return render(request, 'course_detail.html', context)

@login_required
def profile(request):
    """
    Renders the user profile page with enrolled courses.
    """
    enrollments = Enrollment.objects.filter(student=request.user)
    context = {'enrollments': enrollments}
    return render(request, 'profile.html', context)

def user_login(request):
    """
    Handles user login.
    """
    if request.method == 'POST':
        username = request.POST['username']
        password = request.POST['password']
        user = authenticate(request, username=username, password=password)
        if user is not None:
            login(request, user)
            return redirect('index')
        else:
            form = forms.Form()  # Empty form for errors
            return render(request, 'login.html', {'form': form})
    return render(request, 'login.html', {'form': forms.Form()})

def user_logout(request):
    """
    Handles user logout.
    """
    logout(request)
    return redirect('index')

def register(request):
    """
    Handles user registration and sends a welcome email.
    """
    if request.method == 'POST':
        form = RegisterForm(request.POST)
        if form.is_valid():
            user = form.save(commit=False)
            user.set_password(form.cleaned_data['password1'])
            user.save()
            # Send welcome email
            subject = 'Welcome to E-Learn!'
            html_message = render_to_string('emails/welcome_email.html', {
                'user': user,
                'site_url': request.build_absolute_uri('/')
            })
            plain_message = f"Hi {user.username},\n\nWelcome to E-Learn! We're excited to have you on board. Start exploring our courses at {request.build_absolute_uri('/')}\n\nBest,\nThe E-Learn Team"
            try:
                send_mail(
                    subject,
                    plain_message,
                    settings.DEFAULT_FROM_EMAIL,
                    [user.email],
                    html_message=html_message,
                    fail_silently=False,
                )
            except Exception as e:
                messages.warning(request, 'Welcome email could not be sent, but your account was created.')
            login(request, user)
            messages.success(request, 'Registration successful! A welcome email has been sent.')
            return redirect('index')
    else:
        form = RegisterForm()
    return render(request, 'register.html', {'form': form})

@login_required
def enroll(request, slug):
    """
    Handles course enrollment and sends an enrollment email.
    """
    course = get_object_or_404(Course, slug=slug)
    if request.method == 'POST':
        if not Enrollment.objects.filter(student=request.user, course=course).exists():
            enrollment = Enrollment.objects.create(student=request.user, course=course)
            # Send enrollment email
            subject = f'Enrolled in {course.title}'
            html_message = render_to_string('emails/course_enrollment_email.html', {
                'user': request.user,
                'course': course,
                'course_url': request.build_absolute_uri(course.get_absolute_url())
            })
            plain_message = f"Hi {request.user.username},\n\nYou have successfully enrolled in {course.title}! Access the course at {request.build_absolute_uri(course.get_absolute_url())}\n\nBest,\nThe E-Learn Team"
            try:
                send_mail(
                    subject,
                    plain_message,
                    settings.DEFAULT_FROM_EMAIL,
                    [request.user.email],
                    html_message=html_message,
                    fail_silently=False,
                )
            except Exception as e:
                messages.warning(request, 'Enrollment email could not be sent, but you are enrolled.')
            messages.success(request, f'You have successfully enrolled in {course.title}!')
            return redirect('course_detail', slug=course.slug)
    return render(request, 'enroll.html', {'course': course})