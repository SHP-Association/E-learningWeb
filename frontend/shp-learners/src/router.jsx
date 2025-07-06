import React, { useEffect, useState, Suspense } from 'react';

const Home = React.lazy(() => import('./pages/Home.jsx'));
const CourseDetail = React.lazy(() => import('./pages/CourseDetail.jsx'));
const Enroll = React.lazy(() => import('./pages/Enroll.jsx'));
const FAQ = React.lazy(() => import('./pages/FAQ.jsx'));
const Login = React.lazy(() => import('./pages/Login.jsx'));
const Profile = React.lazy(() => import('./pages/Profile.jsx'));
const Register = React.lazy(() => import('./pages/Register.jsx'));
const PasswordReset = React.lazy(() => import('./pages/PasswordReset.jsx'));
const PasswordResetDone = React.lazy(() => import('./pages/PasswordResetDone.jsx'));
const PasswordResetConfirm = React.lazy(() => import('./pages/PasswordResetConfirm.jsx'));
const PasswordResetComplete = React.lazy(() => import('./pages/PasswordResetComplete.jsx'));

const BACKEND_URL = import.meta.env.VITE_APP_BACKEND_URL;

export default function Router({
  currentRoute,
  user,
  courses: initialCourses,
  enrollments: initialEnrollments,
  onLogin,
  navigate,
  addEnrollment
}) {
  const [courses, setCourses] = useState(initialCourses || []);
  const [enrollments, setEnrollments] = useState(initialEnrollments || []);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setLoading(true);
    fetch(`${BACKEND_URL}/api/courses/`, { credentials: 'include' })
      .then(res => {
        if (!res.ok) {
          throw new Error(`HTTP error! status: ${res.status}`);
        }
        return res.json();
      })
      .then(data => setCourses(data))
      .catch(error => {
        console.error("Failed to fetch courses:", error);
      })
      .finally(() => setLoading(false));
  }, []);

  useEffect(() => {
    if (user) {
      setLoading(true);
      fetch(`${BACKEND_URL}/api/enrollments/`, { credentials: 'include' })
        .then(res => {
          if (!res.ok) {
            throw new Error(`HTTP error! status: ${res.status}`);
          }
          return res.json();
        })
        .then(data => setEnrollments(data))
        .catch(error => {
          console.error("Failed to fetch enrollments:", error);
        })
        .finally(() => setLoading(false));
    }
  }, [user]);

  function Redirect({ to }) {
    useEffect(() => {
      navigate(to);
    }, [to, navigate]);
    return null;
  }

  const pathSegments = currentRoute.split('/').filter(Boolean);

  if (pathSegments[0] === 'courses') {
    return <Redirect to="/" />;
  }

  const isHome = pathSegments.length === 0 || pathSegments[0] === 'index';
  const isCourseDetail = pathSegments[0] === 'course' && pathSegments[1];
  const isProfile = pathSegments[0] === 'profile';
  const isLogin = pathSegments[0] === 'login';
  const isRegister = pathSegments[0] === 'register';
  const isEnroll = pathSegments[0] === 'enroll' && pathSegments[1];
  const isFAQ = pathSegments[0] === 'faq';
  const isPasswordReset = pathSegments[0] === 'password_reset';

  let ComponentToRender;
  let componentProps = {};

  if (isHome) {
    ComponentToRender = Home;
    componentProps = { courses, navigate };
  } else if (isCourseDetail) {
    const slug = pathSegments[1];
    const course = courses.find(c => c.slug === slug);
    const isEnrolled = user
      ? (Array.isArray(enrollments) ? enrollments : []).some(
        e => e.student.username === user.username && e.course.slug === slug
      )
      : false;
    ComponentToRender = course ? CourseDetail : () => <p className="text-center text-xl mt-10">Course not found.</p>;
    componentProps = { course, isEnrolled, user, navigate };
  } else if (isProfile) {
    const userEnrollments = user
      ? (Array.isArray(enrollments) ? enrollments : []).filter(
        e => e.student.username === user.username
      )
      : [];
    ComponentToRender = user ? Profile : Login;
    componentProps = user ? { user, enrollments: userEnrollments, navigate } : { onLogin, navigate };
  } else if (isLogin) {
    ComponentToRender = Login;
    componentProps = { onLogin, navigate };
  } else if (isRegister) {
    ComponentToRender = Register;
    componentProps = { navigate };
  } else if (isEnroll) {
    const slug = pathSegments[1];
    const course = courses.find(c => c.slug === slug);
    ComponentToRender = user && course ? Enroll : Login;
    componentProps = user && course ? { course, user, addEnrollment, navigate } : { onLogin, navigate };
  } else if (isFAQ) {
    ComponentToRender = FAQ;
  } else if (isPasswordReset) {
    if (pathSegments[1] === 'done') {
      ComponentToRender = PasswordResetDone;
    } else if (pathSegments[1] && pathSegments[2] && pathSegments[1] !== 'complete') {
      const uid = pathSegments[1];
      const token = pathSegments[2];
      ComponentToRender = PasswordResetConfirm;
      componentProps = { navigate, uid, token };
    } else if (pathSegments[1] === 'complete') {
      ComponentToRender = PasswordResetComplete;
    } else {
      ComponentToRender = PasswordReset;
    }
    componentProps = { ...componentProps, navigate };
  } else {
    ComponentToRender = Home;
    componentProps = { courses, navigate };
  }

  if (loading) {
    return <div className="text-center text-xl mt-10">Loading data...</div>;
  }

  return (
    <Suspense fallback={<div className="text-center text-xl mt-10">Loading component...</div>}>
      <ComponentToRender {...componentProps} />
    </Suspense>
  );
}