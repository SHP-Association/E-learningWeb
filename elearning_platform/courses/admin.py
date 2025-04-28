from django.contrib import admin
from .models import Course, CourseEnrollment, Category, Lesson, Review, Tag, FAQ, Certificate, Profile
from django.contrib.admin import AdminSite

class CustomAdminSite(AdminSite):
    site_header = "E-Learning Platform Admin"
    site_title = "E-Learning Admin Portal"
    index_title = "Welcome to the E-Learning Platform Admin"

admin_site = CustomAdminSite(name='custom_admin')

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

@admin.register(Category)
class CategoryAdmin(admin.ModelAdmin):
    list_display = ['name', 'created_at']
    search_fields = ['name']

@admin.register(Lesson)
class LessonAdmin(admin.ModelAdmin):
    list_display = ['title', 'course', 'order', 'is_preview']
    list_filter = ['course']
    search_fields = ['title', 'course__title']
    ordering = ['order']

@admin.register(Review)
class ReviewAdmin(admin.ModelAdmin):
    list_display = ['course', 'student', 'rating', 'created_at']
    list_filter = ['rating', 'created_at']
    search_fields = ['course__title', 'student__username']

@admin.register(Tag)
class TagAdmin(admin.ModelAdmin):
    list_display = ['name', 'created_at']
    search_fields = ['name']

@admin.register(FAQ)
class FAQAdmin(admin.ModelAdmin):
    list_display = ['course', 'question']
    search_fields = ['course__title', 'question']

@admin.register(Certificate)
class CertificateAdmin(admin.ModelAdmin):
    list_display = ['course', 'student', 'issued_at']
    search_fields = ['course__title', 'student__username']

@admin.register(Profile)
class ProfileAdmin(admin.ModelAdmin):
    list_display = ['user', 'user_type', 'created_at']
    list_filter = ['user_type', 'created_at']
    search_fields = ['user__username', 'user__email', 'bio']
    readonly_fields = ['created_at']
    fieldsets = (
        (None, {
            'fields': ('user', 'user_type', 'bio', 'profile_picture', 'created_at')
        }),
    )

# Register all models with the custom admin site
admin_site.register(Course, CourseAdmin)
admin_site.register(CourseEnrollment, CourseEnrollmentAdmin)
admin_site.register(Category, CategoryAdmin)
admin_site.register(Lesson, LessonAdmin)
admin_site.register(Review, ReviewAdmin)
admin_site.register(Tag, TagAdmin)
admin_site.register(FAQ, FAQAdmin)
admin_site.register(Certificate, CertificateAdmin)
admin_site.register(Profile, ProfileAdmin)