import React, { useState } from 'react';

function PasswordReset({ navigate }) {
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');
  const BACKEND_URL = import.meta.env.VITE_APP_BACKEND_URL;

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMessage('');
    setError('');

    if (!email.trim()) {
      setError('Please enter your email address.');
      return;
    }

    try {
      const response = await fetch(`${BACKEND_URL}/api/password_reset/request/`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email }),
      });

      const data = await response.json();

      if (response.ok) {
        setMessage(data.message || 'If your email is registered, you will receive a password reset link.');
        setTimeout(() => {
          navigate('/password_reset/done');
        }, 1500);
      } else {
        setError(data.error || 'Failed to send password reset email. Please try again.');
      }
    } catch (err) {
      setError('An error occurred. Please try again later.');
    }
  };

  return (
    <div className="flex flex-col items-center justify-center bg-gray-100 px-4">
      <div className="bg-white p-8 rounded-lg shadow-xl border border-gray-200 w-full max-w-md">
        <h1 className="text-4xl font-bold mb-6 text-center text-blue-800">Reset Your Password</h1>
        <p className="text-gray-700 mb-6 text-center">
          Enter your email address below, and we'll send you instructions to reset your password.
        </p>

        {message && (
          <div className="alert alert-success bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded-lg mb-4" role="alert">
            {message}
          </div>
        )}
        {error && (
          <div className="alert alert-error bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg mb-4" role="alert">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="w-full">
          <div className="form-group">
            <label htmlFor="email" className="form-label">Email</label>
            <input
              type="email"
              name="email"
              id="email"
              className="w-full p-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white text-gray-800"
              required
              autoComplete="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>

          <button type="submit" className="btn btn-primary bg-blue-900 text-white hover:bg-blue-700 w-full py-3 mt-4">
            Send Reset Link
          </button>
        </form>

        <p className="mt-6 text-center text-gray-700">
          Remembered your password?{' '}
          <button onClick={() => navigate('/login')} className="text-blue-600 hover:underline">
            Login here
          </button>
        </p>
      </div>
    </div>
  );
}

export default PasswordReset;
