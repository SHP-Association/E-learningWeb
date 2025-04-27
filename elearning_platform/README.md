# E-Learning Platform

Welcome to the E-Learning Platform project! This project is designed to provide an online learning experience with various courses available for users to enhance their skills.

## Project Structure

```
elearning_platform
├── manage.py
├── requirements.txt
├── README.md
├── elearning
│   ├── __init__.py
│   ├── settings.py
│   ├── urls.py
│   ├── wsgi.py
│   └── asgi.py
├── courses
│   ├── __init__.py
│   ├── admin.py
│   ├── apps.py
│   ├── models.py
│   ├── urls.py
│   ├── views.py
│   └── tests.py
├── static
│   ├── css
│   │   └── style.css
│   └── js
│       └── scripts.js
└── templates
    ├── base.html
    ├── courses
    │   ├── course_detail.html
    │   ├── course_list.html
    │   └── home.html
    └── includes
        ├── footer.html
        ├── header.html
        └── modal.html
```

## Setup Instructions

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd elearning_platform
   ```

2. **Create a virtual environment:**
   ```bash
   python -m venv venv
   source venv/bin/activate  # On Windows use `venv\Scripts\activate`
   ```

3. **Install the required packages:**
   ```bash
   pip install -r requirements.txt
   ```

4. **Run the migrations:**
   ```bash
   python manage.py migrate
   ```

5. **Start the development server:**
   ```bash
   python manage.py runserver
   ```

6. **Access the application:**
   Open your web browser and go to `http://127.0.0.1:8000/`.

## Usage

- Navigate through the homepage to explore available courses.
- Click on a course to view its details and enroll.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.