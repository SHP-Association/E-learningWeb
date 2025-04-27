from django.shortcuts import render, get_object_or_404, redirect
from django.db.models import Q
from django.contrib.auth.decorators import login_required
from django.contrib.auth import login, authenticate
from django.contrib import messages
from .models import Course
from .forms import UserRegistrationForm

def home(request):
    featured_courses = Course.objects.filter(is_published=True)[:3]
    return render(request, 'courses/home.html', {'featured_courses': featured_courses})

def course_list(request):
    query = request.GET.get('q', '')
    level = request.GET.get('level', '')
    category = request.GET.get('category', '')
    
    courses = Course.objects.all()
    
    if query:
        courses = courses.filter(
            Q(title__icontains=query) |
            Q(description__icontains=query)
        )
    
    if level:
        courses = courses.filter(level=level)
    
    if category:
        courses = courses.filter(category=category)
        
    context = {
        'courses': courses,
        'query': query,
        'level': level,
        'category': category,
        'levels': Course.LEVEL_CHOICES,
        'categories': Course.objects.values_list('category', flat=True).distinct()
    }
    return render(request, 'courses/course_list.html', context)

def course_detail(request, course_id):
    course = get_object_or_404(Course, id=course_id)
    context = {
        'course': course,
        'instructor': course.instructor,
        'is_enrolled': False
    }
    
    if request.user.is_authenticated:
        # Add enrollment check logic here when you have an Enrollment model
        pass
        
    return render(request, 'courses/course_detail.html', context)

@login_required
def dashboard(request):
    # ...existing code...
    pass

def user_login(request):
    if request.method == 'POST':
        username = request.POST['username']
        password = request.POST['password']
        user = authenticate(request, username=username, password=password)
        
        if user is not None:
            login(request, user)
            return redirect('dashboard')
        else:
            messages.error(request, 'Invalid username or password.')
    
    return render(request, 'accounts/login.html')

def register(request):
    if request.method == 'POST':
        form = UserRegistrationForm(request.POST)
        if form.is_valid():
            user = form.save()
            login(request, user)
            messages.success(request, 'Registration successful.')
            return redirect('dashboard')
    else:
        form = UserRegistrationForm()
    
    return render(request, 'accounts/register.html', {'form': form})