import React, { useEffect, useState } from 'react';
const BACKEND_URL = import.meta.env.VITE_APP_BACKEND_URL;

function Home({ courses: initialCourses, navigate }) {
  const [courses, setCourses] = useState(initialCourses || []);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    setLoading(true);
    fetch(`${BACKEND_URL}/api/courses/`)
      .then(res => {
        if (!res.ok) {
          throw new Error(`HTTP error! status: ${res.status}`);
        }
        return res.json();
      })
      .then(data => {
        setCourses(data);
        setLoading(false);
      })
      .catch(err => {
        console.error("Error fetching courses:", err);
        setError('Failed to load courses. Please try again later.');
        setLoading(false);
      });
  }, []);

  // Helper to render star ratings
  const renderStars = (rating) => {
    const roundedRating = Math.round(rating);
    return (
      <div className="flex items-center">
        {[...Array(5)].map((_, i) => (
          <svg key={i} className={`w-4 h-4 ${i < roundedRating ? 'text-yellow-400' : 'text-gray-300'}`} fill="currentColor" viewBox="0 0 20 20">
            <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.538 1.118l-2.8-2.034a1 1 0 00-1.176 0l-2.8 2.034c-.783.57-1.838-.197-1.538-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.381-1.81.588-1.81h3.462a1 1 0 00.95-.69l1.07-3.292z"></path>
          </svg>
        ))}
      </div>
    );
  };

  return (
    <div className="container mx-auto px-4 py-8 md:py-12">
      {/* Hero Section */}
      <section className="hero relative bg-gradient-to-r from-blue-600 to-indigo-700 text-white py-16 md:py-24 lg:py-32 text-center rounded-3xl shadow-2xl mb-12 md:mb-16 overflow-hidden">
        {/* Animated Background Circles */}
        <div className="absolute inset-0 pointer-events-none z-0">
          <div className="absolute top-[-60px] left-[-60px] w-72 h-72 bg-blue-400 opacity-20 rounded-full animate-pulse-slow"></div>
          <div className="absolute bottom-[-80px] right-[-80px] w-96 h-96 bg-indigo-400 opacity-20 rounded-full animate-pulse-slower"></div>
        </div>
        {/* Background Overlay for Visual Interest */}
        <div className="absolute inset-0 bg-black opacity-10 rounded-3xl z-0"></div>
        <div className="relative z-10">
          <h1 className="text-4xl sm:text-5xl md:text-6xl lg:text-7xl font-extrabold mb-4 md:mb-6 leading-tight drop-shadow-lg animate-fade-in-down">
            Welcome to <span className="text-yellow-300">SHP-Learner</span>
          </h1>
          <p className="text-lg sm:text-xl md:text-2xl lg:text-3xl mb-8 md:mb-10 opacity-90 max-w-xs sm:max-w-md md:max-w-2xl lg:max-w-3xl mx-auto px-4 animate-fade-in">
            Explore our wide range of meticulously crafted courses to elevate your skills and career.
          </p>
          <button
            onClick={() => navigate('/courses')}
            className="btn bg-white text-blue-700 hover:bg-blue-100 hover:text-blue-800 px-8 py-3 md:px-10 md:py-4 text-lg md:text-xl font-semibold rounded-full shadow-xl transform hover:-translate-y-1 transition duration-300 ease-in-out border-2 border-transparent hover:border-blue-700 animate-scale-in"
          >
            Browse All Courses
          </button>
        </div>
      </section>

      {/* Section Divider */}
      <div className="flex items-center justify-center my-12">
        <span className="h-1 w-24 bg-blue-200 rounded-full"></span>
        <span className="mx-4 text-2xl font-bold text-blue-700 tracking-wide">Courses</span>
        <span className="h-1 w-24 bg-blue-200 rounded-full"></span>
      </div>

      {/* Courses List */}
      <h2 className="text-3xl sm:text-4xl lg:text-5xl font-extrabold mb-8 md:mb-12 text-gray-800 text-center relative">
        <span className="relative inline-block pb-2">
          Discover Our Popular Courses
          <span className="absolute bottom-0 left-1/2 w-20 h-1 bg-blue-500 transform -translate-x-1/2 rounded-full"></span>
        </span>
      </h2>

      {loading ? (
        <div className="flex justify-center items-center py-16">
          <div className="loader ease-linear rounded-full border-8 border-t-8 border-gray-200 h-24 w-24"></div>
          <p className="text-center text-xl text-blue-600 ml-4">Loading courses...</p>
        </div>
      ) : error ? (
        <div className="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg text-center mx-auto max-w-md mt-8 shadow-md">
          <p className="text-lg">{error}</p>
          <p className="text-sm mt-2">Please check your internet connection or try again later.</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 xl:grid-cols-4 gap-6 md:gap-8">
          {courses && courses.length > 0 ? (
            courses.map((course, idx) => (
              <div
                key={course.slug}
                className="yt-card group bg-white rounded-xl shadow-md hover:shadow-2xl border border-gray-100 transition-all duration-300 flex flex-col cursor-pointer overflow-hidden"
                onClick={() => navigate(`/course/${course.slug}`)}
              >
                {/* Thumbnail */}
                <div className="relative w-full aspect-video bg-gray-100 overflow-hidden">
                  {course.thumbnail ? (
                    <img
                      src={course.thumbnail}
                      alt={course.title}
                      className="w-full h-full object-contain bg-white group-hover:scale-105 transition-transform duration-300"
                      style={{ maxHeight: '220px' }}
                    />
                  ) : course.category?.image ? (
                    <img
                      src={course.category.image}
                      alt={course.category.name}
                      className="w-full h-full object-cover opacity-80"
                    />
                  ) : (
                    <div className="flex items-center justify-center h-full w-full text-gray-400">
                      <svg className="w-12 h-12" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M12 2C6.48 2 2 6.48...z" />
                      </svg>
                    </div>
                  )}
                  {/* Badges */}
                  {idx === 0 && (
                    <span className="absolute top-2 left-2 bg-yellow-400 text-blue-900 text-xs font-semibold px-2 py-1 rounded z-10 shadow animate-pulse">
                      🌟 FEATURED
                    </span>
                  )}
                  {course.is_free && (
                    <span className="absolute top-2 right-2 bg-green-500 text-white text-xs font-semibold px-2 py-1 rounded z-10 shadow">
                      FREE
                    </span>
                  )}
                </div>
                {/* Card Body */}
                <div className="flex flex-row p-3 gap-3">
                  {/* Instructor Avatar */}
                  <div className="flex-shrink-0">
                    {course.instructor?.profile_picture ? (
                      <img
                        src={course.instructor.profile_picture}
                        alt={course.instructor.username}
                        className="w-10 h-10 rounded-full object-cover border border-gray-200"
                      />
                    ) : (
                      <div className="w-10 h-10 rounded-full bg-blue-100 flex items-center justify-center text-blue-700 font-bold text-lg">
                        {course.instructor?.username?.[0]?.toUpperCase() || "?"}
                      </div>
                    )}
                  </div>
                  {/* Info */}
                  <div className="flex-1 min-w-0">
                    <h3 className="text-base md:text-lg font-bold text-blue-900 mb-1 truncate group-hover:text-blue-700 transition-colors">
                      {course.title}
                    </h3>
                    <div className="flex items-center text-xs text-gray-500 mb-1">
                      {renderStars(course.average_rating)}
                      <span className="ml-1 font-semibold">{Number(course.average_rating || 0).toFixed(1)}</span>
                      <span className="ml-1">({course.number_of_reviews || 0})</span>
                    </div>
                    <div className="flex flex-wrap gap-1 text-xs mb-1">
                      <span className="bg-blue-50 text-blue-700 px-2 py-0.5 rounded">{course.category?.name || 'Uncategorized'}</span>
                      <span className="bg-indigo-50 text-indigo-700 px-2 py-0.5 rounded">{course.level ? course.level.charAt(0).toUpperCase() + course.level.slice(1) : 'N/A'}</span>
                      <span className="bg-gray-50 text-gray-700 px-2 py-0.5 rounded">{course.total_lectures || 0} Lectures</span>
                    </div>
                    <p className="text-xs text-gray-700 mt-1 line-clamp-2">
                      {course.short_description ||
                        (course.description?.split(' ').slice(0, 20).join(' ') +
                          (course.description?.split(' ').length > 20 ? '...' : '')) ||
                        'No description available.'}
                    </p>
                  </div>
                </div>
                {/* Card Footer */}
                <div className="flex items-center justify-between px-3 pb-3">
                  <span className="text-xs text-gray-500">
                    By <span className="font-semibold text-blue-700">{course.instructor?.username || 'N/A'}</span>
                  </span>
                  <span className={`font-bold text-sm ${course.is_free ? 'text-green-600' : 'text-blue-800'}`}>
                    {course.is_free
                      ? 'Free'
                      : `₹${typeof course.price === 'number'
                        ? course.price.toFixed(2)
                        : Number(parseFloat(course.price)).toFixed(2)}`}
                  </span>
                </div>
              </div>
            ))
          ) : (
            <p className="text-gray-600 col-span-full text-center py-8 md:py-12 text-lg md:text-xl italic">
              No courses available at the moment. Please check back soon!
            </p>
          )}
        </div>
      )}
      {/* Tailwind CSS Customizations and Animations */}
      <style>{`
        @keyframes fade-in-down {
          from {
            opacity: 0;
            transform: translateY(-20px);
          }
          to {
            opacity: 1;
            transform: translateY(0);
          }
        }
        @keyframes fade-in {
          from {
            opacity: 0;
          }
          to {
            opacity: 1;
          }
        }
        @keyframes scale-in {
          from {
            opacity: 0;
            transform: scale(0.9);
          }
          to {
            opacity: 1;
            transform: scale(1);
          }
        }
        @keyframes pulse-slow {
          0%, 100% { opacity: 0.18; transform: scale(1);}
          50% { opacity: 0.28; transform: scale(1.08);}
        }
        @keyframes pulse-slower {
          0%, 100% { opacity: 0.13; transform: scale(1);}
          50% { opacity: 0.22; transform: scale(1.12);}
        }
        @keyframes bounce-slow {
          0%, 100% { transform: translateY(0);}
          50% { transform: translateY(-8px);}
        }
        .animate-fade-in-down {
          animation: fade-in-down 1s ease-out forwards;
        }
        .animate-fade-in {
          animation: fade-in 1.5s ease-out forwards;
        }
        .animate-scale-in {
          animation: scale-in 0.7s ease-out forwards;
        }
        .animate-pulse-slow {
          animation: pulse-slow 4s infinite;
        }
        .animate-pulse-slower {
          animation: pulse-slower 7s infinite;
        }
        .animate-bounce-slow {
          animation: bounce-slow 2.5s infinite;
        }
        .loader {
          border-top-color: #3498db;
          -webkit-animation: spinner 1.5s linear infinite;
          animation: spinner 1.5s linear infinite;
        }
        @-webkit-keyframes spinner {
          0% { -webkit-transform: rotate(0deg); }
          100% { -webkit-transform: rotate(360deg); }
        }
        @keyframes spinner {
          0% { transform: rotate(0deg); }
          100% { transform: rotate(360deg); }
        }
        .yt-card {
          transition: box-shadow 0.2s, transform 0.2s;
          background: #fff;
          display: flex;
          flex-direction: column;
          border-radius: 1rem;
        }
        .yt-card .aspect-video {
          aspect-ratio: 16 / 9;
        }
        @media (max-width: 640px) {
          .yt-card .aspect-video {
            min-height: 160px;
          }
        }
        @media (min-width: 641px) and (max-width: 1023px) {
          .yt-card .aspect-video {
            min-height: 180px;
          }
        }
        @media (min-width: 1024px) {
          .yt-card .aspect-video {
            min-height: 200px;
          }
        }
        .yt-card .line-clamp-2 {
          display: -webkit-box;
          -webkit-line-clamp: 2;
          -webkit-box-orient: vertical;
          overflow: hidden;
        }
      `}</style>
    </div>
  );
}

export default Home;