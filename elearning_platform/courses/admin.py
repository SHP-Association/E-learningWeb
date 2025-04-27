from django.contrib import admin
from .models import Course, CourseEnrollment

@admin.register(Course)
class CourseAdmin(admin.ModelAdmin):
    list_display = ['title', 'instructor', 'price', 'level', 'total_enrollments', 'is_active']
    list_filter = ['level', 'category', 'is_active']
    search_fields = ['title', 'description']
    readonly_fields = ['created_at']
    
    def total_enrollments(self, obj):
        return obj.courseenrollment_set.count()
    
    def get_queryset(self, request):
        qs = super().get_queryset(request)
        if not request.user.is_superuser:
            return qs.filter(instructor=request.user)
        return qs
    
    def save_model(self, request, obj, form, change):
        if not obj.instructor:
            obj.instructor = request.user
        super().save_model(request, obj, form, change)

@admin.register(CourseEnrollment)
class CourseEnrollmentAdmin(admin.ModelAdmin):
    list_display = ['course', 'student', 'enrolled_at', 'payment_status']
    list_filter = ['payment_status', 'enrolled_at']
    search_fields = ['course__title', 'student__username']
    
    def get_queryset(self, request):
        qs = super().get_queryset(request)
        if not request.user.is_superuser:
            return qs.filter(course__instructor=request.user)
        return qs