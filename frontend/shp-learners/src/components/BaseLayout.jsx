import React, { useState, useContext, useRef, useEffect } from 'react';
import { UserContext } from '../App.jsx';
import logo from '../assets/logo.png';

function BaseLayout({ children }) {
  const { user, handleLogout, navigate } = useContext(UserContext);
  const [showProfileDropdown, setShowProfileDropdown] = useState(false);
  const profileBtnRef = useRef(null);
  const profileDropdownRef = useRef(null);

  // Close dropdown on outside click
  useEffect(() => {
    function handleClickOutside(event) {
      if (
        profileBtnRef.current &&
        !profileBtnRef.current.contains(event.target) &&
        profileDropdownRef.current &&
        !profileDropdownRef.current.contains(event.target)
      ) {
        setShowProfileDropdown(false);
      }
    }
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  // Only one profile dropdown for all screens
  return (
    <div className="min-h-screen flex flex-col bg-gray-100 font-sans">
      <nav className="bg-blue-900 p-4 shadow-lg">
        <div className="container mx-auto flex justify-between items-center">
          {/* Logo and Site Title */}
          <button
            className="flex items-center text-white text-xl md:text-2xl font-bold rounded-lg px-2 py-1 md:px-3 md:py-2 transition-colors duration-300 hover:bg-blue-800 focus:outline-none"
            onClick={() => navigate('/')}
            aria-label="Go to homepage"
          >
            <img src={logo} alt="SHP-Learner Logo" className="w-8 h-8 md:w-10 md:h-10 mr-2 rounded-full" />
            SHP-Learner
          </button>

          {/* Navigation */}
          <div className="flex items-center space-x-4 md:space-x-8">
            <button
              onClick={() => navigate('/')}
              className="text-white hover:text-blue-200 transition-colors duration-300 rounded-lg hover:bg-blue-800 text-base md:text-lg px-4 py-2 focus:outline-none hidden md:block"
            >
              Home
            </button>
            {user && user.is_staff && (
              <button
                onClick={() => window.open(`${import.meta.env.VITE_APP_BACKEND_URL}/admin`, '_blank')}
                className="text-white hover:text-blue-200 transition-colors duration-300 rounded-lg hover:bg-blue-800 text-base md:text-lg px-4 py-2 focus:outline-none hidden md:block"
              >
                Admin
              </button>
            )}
            {!user && (
              <button
                onClick={() => navigate('/login')}
                className="text-white hover:text-blue-200 transition-colors duration-300 rounded-lg hover:bg-blue-800 text-base md:text-lg px-4 py-2 focus:outline-none hidden md:block"
              >
                Login
              </button>
            )}
            {/* Profile/Login button for all screens */}
            {user ? (
              <div className="relative" ref={profileBtnRef}>
                <button
                  onClick={() => setShowProfileDropdown((v) => !v)}
                  className="focus:outline-none"
                  aria-label="Open profile menu"
                  style={{ padding: 0, background: 'none', border: 'none' }}
                >
                  {user.profile_picture ? (
                    <img
                      src={user.profile_picture}
                      alt="Profile"
                      className="w-9 md:w-11 h-9 md:h-11 rounded-full border-2 border-yellow-400 object-cover shadow"
                    />
                  ) : (
                    <span className="flex items-center justify-center bg-yellow-400 rounded-full w-9 h-9 md:w-11 md:h-11">
                      <svg xmlns="http://www.w3.org/2000/svg" className="w-6 md:w-7 h-6 md:h-7 text-blue-900" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <circle cx="12" cy="8" r="4" stroke="currentColor" strokeWidth="2" fill="none"/>
                        <path stroke="currentColor" strokeWidth="2" d="M4 20c0-4 4-7 8-7s8 3 8 7"/>
                      </svg>
                    </span>
                  )}
                </button>
                {showProfileDropdown && (
                  <div
                    ref={profileDropdownRef}
                    className="absolute right-0 mt-2 w-44 bg-white rounded-lg shadow-lg z-50 border border-gray-200"
                  >
                    <button
                      onClick={() => { navigate('/profile'); setShowProfileDropdown(false); }}
                      className="block w-full text-left px-4 py-2 text-gray-800 hover:bg-blue-100 hover:text-blue-900 rounded-t-lg transition-colors duration-200"
                    >
                      Profile
                    </button>
                    <button
                      onClick={() => { handleLogout(); setShowProfileDropdown(false); }}
                      className="block w-full text-left px-4 py-2 text-gray-800 hover:bg-blue-100 hover:text-blue-900 rounded-b-lg transition-colors duration-200"
                    >
                      Logout
                    </button>
                  </div>
                )}
              </div>
            ) : (
              <button
                onClick={() => navigate('/login')}
                className="text-white bg-yellow-400 hover:bg-yellow-500 rounded-full px-5 py-2 font-semibold text-base shadow focus:outline-none md:hidden"
              >
                Login
              </button>
            )}
          </div>
        </div>
      </nav>

      <main className="container mx-auto p-4 md:p-6 flex-grow">
        {children}
      </main>

      <footer className="bg-blue-900 text-white py-10 mt-auto shadow-inner">
        <div className="footer-container container mx-auto grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-10 px-4">
          <div className="footer-about">
            <h4 className="text-2xl font-bold mb-4 flex items-center gap-2">
              <span className="bg-yellow-400 text-blue-900 px-2 py-1 rounded-lg font-black">SHP</span>
              About Us
            </h4>
            <p className="text-gray-300 text-base leading-relaxed">
              SHP-Learnering Platform is dedicated to providing high-quality courses to help you enhance your skills and achieve your career goals. A product by SHP Association.<br />
              <span className="inline-block mt-2 font-semibold text-yellow-300">Join us to start learning today!</span>
            </p>
          </div>

          <div className="footer-links">
            <h4 className="text-2xl font-bold mb-4">Quick Links</h4>
            <ul>
              <li className="mb-2">
                <button onClick={() => navigate('/')} className="text-gray-300 hover:text-yellow-300 transition-colors duration-300 text-base font-medium focus:outline-none">Home</button>
              </li>
              <li className="mb-2">
                <a href="https://wa.me/9399613606" target="_blank" rel="noopener noreferrer" className="text-gray-300 hover:text-yellow-300 transition-colors duration-300 text-base font-medium">Contact Us</a>
              </li>
              <li className="mb-2">
                <button onClick={() => navigate('/faq')} className="text-gray-300 hover:text-yellow-300 transition-colors duration-300 text-base font-medium focus:outline-none">FAQ</button>
              </li>
            </ul>
          </div>

          <div className="footer-contact">
            <h4 className="text-2xl font-bold mb-4">Contact Us</h4>
            <div className="space-y-2">
              <p className="text-gray-300 text-base flex items-center">
                <i className="fas fa-envelope mr-2 text-yellow-300"></i>
                <a href="mailto:sandeshpatel.sp.93@gmail.com" className="hover:text-yellow-300 transition-colors duration-300">sandeshpatel.sp.93@gmail.com</a>
              </p>
              <p className="text-gray-300 text-base flex items-center">
                <i className="fas fa-phone-alt mr-2 text-yellow-300"></i>
                <a href="tel:+919399613606" className="hover:text-yellow-300 transition-colors duration-300">+91 9399613606</a>
              </p>
              <p className="text-gray-300 text-base flex items-center">
                <i className="fas fa-map-marker-alt mr-2 text-yellow-300"></i>
                420 Kareli St, MP Narsinghpur, India
              </p>
            </div>
          </div>

          <div className="footer-social">
            <h4 className="text-2xl font-bold mb-4">Follow Us</h4>
            <div className="social-icons flex space-x-5 mt-2">
              <a href="https://x.com/SandeshPat007?t=teYEP7w9aNZYSYKc0sF7dQ&s=09" className="text-gray-300 hover:text-yellow-300 transition-colors duration-300 text-3xl" target="_blank" rel="noopener noreferrer" aria-label="Follow us on X"><i className="fab fa-twitter"></i></a>
              <a href="https://www.linkedin.com/in/sandesh-patel07" className="text-gray-300 hover:text-yellow-300 transition-colors duration-300 text-3xl" target="_blank" rel="noopener noreferrer" aria-label="Follow us on LinkedIn"><i className="fab fa-linkedin-in"></i></a>
              <a href="https://www.instagram.com/sandesh_patel007" className="text-gray-300 hover:text-yellow-300 transition-colors duration-300 text-3xl" target="_blank" rel="noopener noreferrer" aria-label="Follow us on Instagram"><i className="fab fa-instagram"></i></a>
            </div>
            <div className="mt-4">
              <span className="inline-block bg-yellow-400 text-blue-900 px-3 py-1 rounded-full font-bold text-xs">#KeepLearning</span>
            </div>
          </div>
        </div>

        <div className="footer-bottom text-center text-gray-400 mt-10 border-t border-blue-800 pt-5 px-4">
          <p className="text-base">
            &copy; {new Date().getFullYear()} <span className="font-bold text-yellow-300">SHP-Learnering Platform</span>. All Rights Reserved |
            <button onClick={() => alert('Privacy Policy clicked!')} className="text-gray-400 hover:text-yellow-300 transition-colors duration-300 ml-2 focus:outline-none">Privacy Policy</button> |
            <button onClick={() => alert('Terms of Service clicked!')} className="text-gray-400 hover:text-yellow-300 transition-colors duration-300 ml-2 focus:outline-none">Terms of Service</button>
          </p>
        </div>
      </footer>
    </div>
  );
}

export default BaseLayout;