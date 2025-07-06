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
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8 md:gap-10 lg:gap-12">
          {courses && courses.length > 0 ? (
            courses.map((course, idx) => (
              <div
  key={course.slug}
  className="card bg-white p-6 rounded-3xl shadow-md hover:shadow-2xl transition-all duration-300 ease-in-out border border-gray-100 transform hover:-translate-y-2 hover:scale-[1.03] relative overflow-hidden group flex flex-col cursor-pointer"
  onClick={() => navigate(`/course/${course.slug}`)}
  style={{ minHeight: 440 }}
>
  {/* FEATURED Badge */}
  {idx === 0 && (
    <span className="absolute top-4 left-4 bg-yellow-400 text-blue-900 text-xs font-semibold px-3 py-1 rounded-full z-10 tracking-wide shadow-md animate-pulse">
      🌟 FEATURED
    </span>
  )}

  {/* FREE Badge */}
  {course.is_free && (
    <span className="absolute top-4 right-4 bg-green-500 text-white text-xs font-semibold px-3 py-1 rounded-full z-10 shadow-md">
      FREE
    </span>
  )}

  {/* Thumbnail */}
  <div className="w-full h-48 overflow-hidden rounded-xl mb-4 bg-gray-100 flex items-center justify-center">
    {course.thumbnail ? (
      <img
        src={course.thumbnail}
        alt={course.title}
        className="w-full h-full object-cover transition-transform duration-300 ease-in-out group-hover:scale-105"
      />
    ) : course.category?.image ? (
      <img
        src={course.category.image}
        alt={course.category.name}
        className="w-full h-full object-cover opacity-80"
      />
    ) : (
      <div className="text-gray-400 flex flex-col items-center justify-center h-full w-full">
        <svg className="w-12 h-12 mb-1 text-gray-300" fill="currentColor" viewBox="0 0 24 24">
          <path d="M12 2C6.48 2 2 6.48...z" />
        </svg>
        <span className="text-sm">No Image Available</span>
      </div>
    )}
  </div>

  {/* Title */}
  <h3 className="text-xl font-bold mb-1 text-blue-800 group-hover:text-blue-900 transition-colors duration-300 leading-snug">
    {course.title}
  </h3>

  {/* Rating */}
  <div className="flex items-center text-sm text-gray-600 mb-3">
    {renderStars(course.average_rating)}
    <span className="ml-2 font-semibold">{Number(course.average_rating || 0).toFixed(1)}</span>
    <span className="ml-1 text-xs text-gray-500">({course.number_of_reviews || 0} reviews)</span>
  </div>

  {/* Details Info */}
  <div className="text-sm text-gray-600 mb-4 space-y-1">
    <p>
      <span className="font-medium">📂 Category:</span>{" "}
      <span className="text-blue-600 font-semibold">{course.category?.name || 'Uncategorized'}</span>
    </p>
    <p>
      <span className="font-medium">🎯 Level:</span>{" "}
      {course.level ? course.level.charAt(0).toUpperCase() + course.level.slice(1) : 'N/A'}
    </p>
    <p>
      <span className="font-medium">🎥 Lectures:</span> {course.total_lectures || 0}
    </p>
  </div>

  {/* Description */}
  <p className="text-sm text-gray-700 mb-4 leading-relaxed line-clamp-3">
    {course.short_description ||
      (course.description?.split(' ').slice(0, 25).join(' ') +
        (course.description.split(' ').length > 25 ? '...' : '')) ||
      'No description available.'}
  </p>

  {/* Instructor Info */}
  <div className="flex items-center text-sm text-gray-700 mb-4">
    {course.instructor?.profile_picture ? (
      <img
        src={course.instructor.profile_picture}
        alt={course.instructor.username}
        className="w-8 h-8 rounded-full object-cover mr-2 border border-gray-300"
      />
    ) : (
      <svg className="w-8 h-8 mr-2 text-gray-400" fill="currentColor" viewBox="0 0 20 20">
        <path d="M10 9a3 3 0 100-6...z" />
      </svg>
    )}
    <span>
      <span className="font-semibold">By:</span>{" "}
      <span className="ml-1 text-blue-600">{course.instructor?.username || 'N/A'}</span>
    </span>
  </div>

  {/* Price */}
  <div className="flex items-center justify-between text-gray-800 border-t pt-4">
    <strong className="text-base">Price:</strong>
    {course.is_free ? (
      <span className="text-green-600 font-bold text-lg">Free</span>
    ) : (
      <span className="text-blue-800 font-bold text-lg">
        ₹{typeof course.price === 'number'
          ? course.price.toFixed(2)
          : Number(parseFloat(course.price)).toFixed(2)}
      </span>
    )}
  </div>

  {/* CTA Button */}
  <button
    onClick={(e) => {
      e.stopPropagation();
      navigate(`/course/${course.slug}`);
    }}
    className="btn bg-blue-800 text-white mt-5 py-2 rounded-lg w-full hover:bg-blue-700 transition-transform hover:scale-105 text-sm font-semibold"
  >
    View Details
  </button>
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
      `}</style>
    </div>
  );
}

export default Home;