document.addEventListener('DOMContentLoaded', function() {
    // Tab switching
    const navItems = document.querySelectorAll('.nav-item');
    const tabContents = document.querySelectorAll('.tab-content');

    navItems.forEach(item => {
        item.addEventListener('click', (e) => {
            e.preventDefault();
            const targetTab = item.dataset.tab;
            
            // Update active states
            navItems.forEach(nav => nav.classList.remove('active'));
            tabContents.forEach(tab => tab.classList.remove('active'));
            
            item.classList.add('active');
            document.getElementById(targetTab).classList.add('active');
        });
    });

    // Course actions
    document.getElementById('addCourseBtn').addEventListener('click', () => {
        window.location.href = '/courses/create/';
    });

    // Edit course
    document.querySelectorAll('.btn-edit').forEach(btn => {
        btn.addEventListener('click', () => {
            const courseId = btn.dataset.id;
            window.location.href = `/courses/${courseId}/edit/`;
        });
    });

    // Delete course
    document.querySelectorAll('.btn-delete').forEach(btn => {
        btn.addEventListener('click', async () => {
            if(confirm('Are you sure you want to delete this course?')) {
                const courseId = btn.dataset.id;
                try {
                    const response = await fetch(`/api/courses/${courseId}/`, {
                        method: 'DELETE',
                        headers: {
                            'X-CSRFToken': getCookie('csrftoken')
                        }
                    });
                    if(response.ok) {
                        btn.closest('.course-card').remove();
                    }
                } catch(error) {
                    console.error('Error:', error);
                }
            }
        });
    });
});

// Helper function to get CSRF token
function getCookie(name) {
    let cookieValue = null;
    if (document.cookie && document.cookie !== '') {
        const cookies = document.cookie.split(';');
        for (let i = 0; i < cookies.length; i++) {
            const cookie = cookies[i].trim();
            if (cookie.substring(0, name.length + 1) === (name + '=')) {
                cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
                break;
            }
        }
    }
    return cookieValue;
}
