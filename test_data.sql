-- ========================================
-- E-Learning Platform Test Data
-- ========================================

-- 1. CATEGORIES (no created_at/updated_at fields)
INSERT INTO "Category_category" (name, slug, description)
VALUES 
    ('Web Development', 'web-development', 'Learn modern web development technologies'),
    ('Data Science', 'data-science', 'Master data analysis and machine learning'),
    ('Mobile Development', 'mobile-development', 'Build iOS and Android applications'),
    ('DevOps', 'devops', 'Learn deployment and infrastructure management');

-- 2. COURSES
INSERT INTO "courses_course" (
    title, slug, short_description, description, what_you_will_learn, requirements, target_audience,
    instructor_id, category_id, price, is_free, is_published, level, duration, 
    total_lectures, average_rating, number_of_reviews, created_at, updated_at
)
VALUES 
    ('Complete Python Bootcamp', 'complete-python-bootcamp', 
     'Master Python from beginner to advanced level',
     'This comprehensive Python course takes you from zero to hero.',
     'Python fundamentals, Data structures, OOP',
     'No prior programming experience required',
     'Beginners who want to learn programming',
     1, 1, 99.99, false, true, 'beginner', '12 weeks', 
     45, 4.5, 120, NOW(), NOW()),
    ('React & Redux Masterclass', 'react-redux-masterclass', 
     'Build modern web applications with React',
     'Learn to build scalable web applications using React and Redux.',
     'React components, State management with Redux',
     'Basic JavaScript knowledge',
     'Web developers wanting to learn React',
     1, 1, 149.99, false, true, 'intermediate', '10 weeks', 
     38, 4.7, 85, NOW(), NOW()),
    ('Machine Learning A-Z', 'machine-learning-a-z', 
     'Comprehensive ML course with hands-on projects',
     'Master machine learning algorithms and implement them in Python.',
     'Supervised learning, Neural networks',
     'Python programming basics',
     'Data scientists, Python developers',
     1, 2, 199.99, false, true, 'advanced', '16 weeks', 
     52, 4.8, 200, NOW(), NOW());

-- 3. LESSONS
INSERT INTO "Lesson_lesson" (course_id, title, slug, content, video_url, "order", is_preview, created_at, updated_at)
VALUES 
    (1, 'Introduction to Python', 'introduction-to-python', 'Welcome to Python!', 'https://example.com/video1', 1, true, NOW(), NOW()),
    (1, 'Variables and Data Types', 'variables-and-data-types', 'Learn about Python variables.', 'https://example.com/video2', 2, false, NOW(), NOW()),
    (1, 'Control Flow', 'control-flow', 'Master if statements and loops.', 'https://example.com/video3', 3, false, NOW(), NOW()),
    (2, 'React Fundamentals', 'react-fundamentals', 'Learn React components and JSX.', 'https://example.com/video4', 1, true, NOW(), NOW()),
    (2, 'State Management', 'state-management', 'Understanding state and props.', 'https://example.com/video5', 2, false, NOW(), NOW()),
    (2, 'Redux Introduction', 'redux-introduction', 'Learn Redux for state management.', 'https://example.com/video6', 3, false, NOW(), NOW()),
    (3, 'ML Basics', 'ml-basics', 'Introduction to machine learning.', 'https://example.com/video7', 1, true, NOW(), NOW()),
    (3, 'Linear Regression', 'linear-regression', 'Understanding regression models.', 'https://example.com/video8', 2, false, NOW(), NOW());

-- 4. QUIZZES
INSERT INTO "Quiz_quiz" (lesson_id, title, description, passing_score, created_at, updated_at)
VALUES 
    (1, 'Python Basics Quiz', 'Test your Python fundamentals', 70, NOW(), NOW()),
    (2, 'Data Types Quiz', 'Quiz on Python data types', 75, NOW(), NOW()),
    (4, 'React Fundamentals Quiz', 'Test your React knowledge', 80, NOW(), NOW()),
    (7, 'ML Basics Quiz', 'Machine learning concepts quiz', 70, NOW(), NOW());

-- 5. QUESTIONS
INSERT INTO "Question_question" (quiz_id, question_text, question_type, "order")
VALUES 
    (1, 'What is Python?', 'mcq', 1),
    (1, 'Which keyword is used to define a function in Python?', 'mcq', 2),
    (2, 'Which of the following is a mutable data type?', 'mcq', 1),
    (3, 'What does JSX stand for?', 'mcq', 1);

-- 6. ANSWER CHOICES
INSERT INTO "Question_answerchoice" (question_id, choice_text, is_correct)
VALUES 
    (1, 'A programming language', true), (1, 'A database', false), (1, 'An operating system', false),
    (2, 'def', true), (2, 'function', false), (2, 'func', false),
    (3, 'List', true), (3, 'Tuple', false), (3, 'String', false),
    (4, 'JavaScript XML', true), (4, 'Java Syntax Extension', false), (4, 'JavaScript Extension', false);

-- 7. FAQs
INSERT INTO "FAQ_faq" (question, answer, category_id, is_published, "order", created_at, updated_at)
VALUES 
    ('How do I enroll in a course?', 'Click on the course and click Enroll button.', 1, true, 1, NOW(), NOW()),
    ('Can I get a refund?', 'Yes, we offer a 30-day money-back guarantee.', 1, true, 2, NOW(), NOW()),
    ('Do I get a certificate?', 'Yes, after completing all lessons and quizzes.', 1, true, 3, NOW(), NOW()),
    ('How long do I have access?', 'Lifetime access to all enrolled courses.', 1, true, 4, NOW(), NOW()),
    ('Are there prerequisites?', 'Check the course description for requirements.', 1, true, 5, NOW(), NOW());

-- VERIFICATION
SELECT 'Categories:' as table_name, COUNT(*) as count FROM "Category_category"
UNION ALL SELECT 'Courses:', COUNT(*) FROM "courses_course"
UNION ALL SELECT 'Lessons:', COUNT(*) FROM "Lesson_lesson"
UNION ALL SELECT 'Quizzes:', COUNT(*) FROM "Quiz_quiz"
UNION ALL SELECT 'Questions:', COUNT(*) FROM "Question_question"
UNION ALL SELECT 'Answers:', COUNT(*) FROM "Question_answerchoice"
UNION ALL SELECT 'FAQs:', COUNT(*) FROM "FAQ_faq";

-- Show inserted data
SELECT id, title, level, price FROM "courses_course";

VALUES 
    ('Web Development', 'Learn modern web development technologies', NOW(), NOW()),
    ('Data Science', 'Master data analysis and machine learning', NOW(), NOW()),
    ('Mobile Development', 'Build iOS and Android applications', NOW(), NOW()),
    ('DevOps', 'Learn deployment and infrastructure management', NOW(), NOW());

-- ========================================
-- 2. COURSES (using admin user id=1 as instructor)
-- ========================================
INSERT INTO "courses_course" (
    title, slug, short_description, description, what_you_will_learn, requirements, target_audience,
    instructor_id, category_id, price, is_free, is_published, level, duration, 
    total_lectures, average_rating, number_of_reviews, created_at, updated_at
)
VALUES 
    ('Complete Python Bootcamp', 'complete-python-bootcamp', 
     'Master Python from beginner to advanced level',
     'This comprehensive Python course takes you from zero to hero. Learn Python basics, data structures, OOP, and more.',
     'Python fundamentals, Data structures, Object-oriented programming, File handling, Error handling',
     'No prior programming experience required. Just bring your enthusiasm!',
     'Beginners who want to learn programming, Developers switching to Python',
     1, 1, 99.99, false, true, 'beginner', '12 weeks', 
     45, 4.5, 120, NOW(), NOW()),
    
    ('React & Redux Masterclass', 'react-redux-masterclass', 
     'Build modern web applications with React and Redux',
     'Learn to build scalable web applications using React, Redux, and modern JavaScript.',
     'React components and hooks, State management with Redux, API integration, Testing',
     'Basic JavaScript knowledge, HTML and CSS fundamentals',
     'Web developers wanting to learn React, JavaScript developers',
     1, 1, 149.99, false, true, 'intermediate', '10 weeks', 
     38, 4.7, 85, NOW(), NOW()),
    
    ('Machine Learning A-Z', 'machine-learning-a-z', 
     'Comprehensive ML course with hands-on projects',
     'Master machine learning algorithms and implement them in Python with real-world projects.',
     'Supervised and unsupervised learning, Neural networks, Model evaluation, Feature engineering',
     'Python programming basics, Basic statistics knowledge',
     'Data scientists, Python developers, AI enthusiasts',
     1, 2, 199.99, false, true, 'advanced', '16 weeks', 
     52, 4.8, 200, NOW(), NOW());

-- ========================================
-- 3. LESSONS
-- ========================================
INSERT INTO "Lesson_lesson" (
    course_id, title, slug, content, video_url, "order", is_preview, created_at, updated_at
)
VALUES 
    (1, 'Introduction to Python', 'introduction-to-python', 
     'Welcome to Python! In this lesson, we cover Python basics and setup your development environment.', 
     'https://example.com/video1', 1, true, NOW(), NOW()),
    (1, 'Variables and Data Types', 'variables-and-data-types',
     'Learn about Python variables, strings, numbers, and basic data types.', 
     'https://example.com/video2', 2, false, NOW(), NOW()),
    (1, 'Control Flow', 'control-flow',
     'Master if statements, loops, and control flow in Python.', 
     'https://example.com/video3', 3, false, NOW(), NOW()),
    (2, 'React Fundamentals', 'react-fundamentals',
     'Learn React components, JSX, and the virtual DOM.', 
     'https://example.com/video4', 1, true, NOW(), NOW()),
    (2, 'State Management', 'state-management',
     'Understanding state, props, and component lifecycle.', 
     'https://example.com/video5', 2, false, NOW(), NOW()),
    (2, 'Redux Introduction', 'redux-introduction',
     'Learn Redux for state management in large applications.', 
     'https://example.com/video6', 3, false, NOW(), NOW()),
    (3, 'ML Basics', 'ml-basics',
     'Introduction to machine learning concepts and terminology.', 
     'https://example.com/video7', 1, true, NOW(), NOW()),
    (3, 'Linear Regression', 'linear-regression',
     'Understanding and implementing linear regression models.', 
     'https://example.com/video8', 2, false, NOW(), NOW());

-- ========================================
-- 4. QUIZZES
-- ========================================
INSERT INTO "Quiz_quiz" (lesson_id, title, description, passing_score, created_at, updated_at)
VALUES 
    (1, 'Python Basics Quiz', 'Test your understanding of Python fundamentals', 70, NOW(), NOW()),
    (2, 'Data Types Quiz', 'Quiz on Python data types and variables', 75, NOW(), NOW()),
    (4, 'React Fundamentals Quiz', 'Test your React knowledge', 80, NOW(), NOW()),
    (7, 'ML Basics Quiz', 'Machine learning concepts quiz', 70, NOW(), NOW());

-- ========================================
-- 5. QUESTIONS & ANSWER CHOICES
-- ========================================
INSERT INTO "Question_question" (quiz_id, question_text, question_type, "order")
VALUES 
    (1, 'What is Python?', 'mcq', 1),
    (1, 'Which keyword is used to define a function in Python?', 'mcq', 2),
    (2, 'Which of the following is a mutable data type?', 'mcq', 1),
    (3, 'What does JSX stand for?', 'mcq', 1);

INSERT INTO "Question_answerchoice" (question_id, choice_text, is_correct)
VALUES 
    (1, 'A programming language', true),
    (1, 'A database', false),
    (1, 'An operating system', false),
    (2, 'def', true),
    (2, 'function', false),
    (2, 'func', false),
    (3, 'List', true),
    (3, 'Tuple', false),
    (3, 'String', false),
    (4, 'JavaScript XML', true),
    (4, 'Java Syntax Extension', false),
    (4, 'JavaScript Extension', false);

-- ========================================
-- 6. FAQs
-- ========================================
INSERT INTO "FAQ_faq" (question, answer, category_id, is_published, "order", created_at, updated_at)
VALUES 
    ('How do I enroll in a course?', 'Click on the course you want and click the "Enroll" button. You will need to create an account first.', 1, true, 1, NOW(), NOW()),
    ('Can I get a refund?', 'Yes, we offer a 30-day money-back guarantee for all courses.', 1, true, 2, NOW(), NOW()),
    ('Do I get a certificate after completion?', 'Yes, you will receive a certificate of completion once you finish all lessons and quizzes.', 1, true, 3, NOW(), NOW()),
    ('How long do I have access to the course?', 'You have lifetime access to all enrolled courses.', 1, true, 4, NOW(), NOW()),
    ('Are there any prerequisites?', 'Prerequisites vary by course. Check the course description for specific requirements.', 1, true, 5, NOW(), NOW());

-- VERIFICATION
SELECT COUNT(*) as categories FROM "Category_category";
SELECT COUNT(*) as courses FROM "courses_course";
SELECT COUNT(*) as lessons FROM "Lesson_lesson";
SELECT COUNT(*) as quizzes FROM "Quiz_quiz";
SELECT COUNT(*) as questions FROM "Question_question";
SELECT COUNT(*) as faqs FROM "FAQ_faq";
