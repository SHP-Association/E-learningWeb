from django.shortcuts import render, get_object_or_404, redirect
from django.db.models import Q
from django.contrib.auth.decorators import login_required
from django.contrib.auth import login, authenticate
from django.contrib import messages
from django.conf import settings
from django.templatetags.static import static
from .models import Course
from .forms import UserRegistrationForm, CourseForm, ProfileForm

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
    thumbnail_url = course.thumbnail.url if course.thumbnail else static('images/default-course-thumbnail.png')
    context = {
        'course': course,
        'instructor': course.instructor,
        'thumbnail_url': thumbnail_url,
        'is_enrolled': False,
        'google_form_url': settings.GOOGLE_FORM_URL,
    }
    
    if request.user.is_authenticated:
        # Add enrollment check logic here when you have an Enrollment model
        pass
        
    return render(request, 'courses/course_detail.html', context)

# @login_required
def dashboard(request):
    # profile_picture_url = request.user.profile.profile_picture.url if request.user.profile.profile_picture 
    # context = {
    #     'profile_picture_url': profile_picture_url,
    #     # ...other context data...
    # }
    return render(request, 'dashboard/provider_dashboard.html')

@login_required
def create_course(request):
    if request.method == 'POST':
        form = CourseForm(request.POST, request.FILES)
        if form.is_valid():
            course = form.save(commit=False)
            course.instructor = request.user
            course.save()
            messages.success(request, 'Course created successfully.')
            return redirect('dashboard')
        else:
            messages.error(request, 'Error creating course. Please check the form.')
    else:
        form = CourseForm()
    
    return render(request, 'courses/create_course.html', {'form': form})

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
        user_form = UserRegistrationForm(request.POST)
        profile_form = ProfileForm(request.POST, request.FILES)
        if user_form.is_valid() and profile_form.is_valid():
            user = user_form.save(commit=False)
            user.set_password(user_form.cleaned_data['password'])
            user.save()
            
            # Update the profile created by the signal
            profile = user.profile
            profile.user_type = profile_form.cleaned_data['user_type']
            profile.bio = profile_form.cleaned_data['bio']
            profile.profile_picture = profile_form.cleaned_data['profile_picture']
            profile.save()
            
            login(request, user)
            messages.success(request, 'Registration successful.')
            return redirect('dashboard')
        else:
            messages.error(request, 'Please correct the errors below.')
    else:
        user_form = UserRegistrationForm()
        profile_form = ProfileForm()

    return render(request, 'accounts/register.html', {
        'user_form': user_form,
        'profile_form': profile_form,
    })