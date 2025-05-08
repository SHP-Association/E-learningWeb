from django.shortcuts import render, get_object_or_404, redirect
from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.decorators import login_required
from django.contrib import messages
from .models import Course, Enrollment, CustomUser
from django import forms

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
    Handles user registration.
    """
    if request.method == 'POST':
        form = RegisterForm(request.POST)
        if form.is_valid():
            user = form.save(commit=False)
            user.set_password(form.cleaned_data['password1'])
            user.save()
            login(request, user)
            return redirect('index')
    else:
        form = RegisterForm()
    return render(request, 'register.html', {'form': form})

@login_required
def enroll(request, slug):
    """
    Handles course enrollment.
    """
    course = get_object_or_404(Course, slug=slug)
    if request.method == 'POST':
        if not Enrollment.objects.filter(student=request.user, course=course).exists():
            Enrollment.objects.create(student=request.user, course=course)
            messages.success(request, f'You have successfully enrolled in {course.title}!')
            return redirect('course_detail', slug=course.slug)
    return render(request, 'enroll.html', {'course': course})