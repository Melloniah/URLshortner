import React, { useState } from 'react';
import axios from 'axios';

const API_URL = "https://urlshortner-production-fbd3.up.railway.app";

function App() {
  const [url, setUrl] = useState('');
  const [shortUrl, setShortUrl] = useState('');
  const [shortCode, setShortCode] = useState('');
  const [originalUrl, setOriginalUrl] = useState('');
  const [clicks, setClicks] = useState(0);
  const [createdAt, setCreatedAt] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [copied, setCopied] = useState(false);
  const [checkingStats, setCheckingStats] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setShortUrl('');
    setLoading(true);

    try {
      const response = await axios.post(`${API_URL}/api/shorten`, { url });
      setShortUrl(response.data.short_url);
      setShortCode(response.data.short_code);
      setOriginalUrl(response.data.long_url);
      setClicks(0);
      setCreatedAt(new Date().toLocaleString());
      setLoading(false);
    } catch (err) {
      setError('Failed to shorten URL. Please check your URL and try again.');
      setLoading(false);
    }
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(shortUrl);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  const refreshStats = async () => {
    if (!shortCode) return;
    setCheckingStats(true);
    try {
      const response = await axios.get(`${API_URL}/api/stats/${shortCode}`);
      setClicks(response.data.clicks);
      setCheckingStats(false);
    } catch (err) {
      console.error('Failed to fetch stats');
      setCheckingStats(false);
    }
  };

  return (
    // Solid Deep Green Background
    <div className="min-h-screen bg-emerald-900 flex items-center justify-center p-4 font-sans">
      <div className="w-full max-w-md">
        <div className="bg-white rounded-2xl shadow-xl p-8">
          <div className="text-center mb-8">
            {/* Solid Deep Green Icon Container */}
            <div className="inline-block p-3 bg-emerald-700 rounded-xl mb-4">
              <svg className="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
              </svg>
            </div>
            <h1 className="text-3xl font-bold text-emerald-900 mb-1">GoLink</h1>
            <p className="text-gray-500 text-sm uppercase tracking-wide">URL Shortener Service</p>
          </div>

          <form onSubmit={handleSubmit} className="space-y-6">
            <div>
              <label htmlFor="url" className="block text-xs font-bold text-emerald-900 uppercase mb-2">
                Enter Long URL
              </label>
              <input
                type="url"
                id="url"
                value={url}
                onChange={(e) => setUrl(e.target.value)}
                placeholder="https://example.com/very-long-url"
                required
                disabled={loading}
                className="w-full px-4 py-3 border-2 border-gray-200 rounded-lg focus:border-emerald-600 outline-none transition-all disabled:bg-gray-50 text-gray-700"
              />
            </div>

            <button
              type="submit"
              disabled={loading}
              className="w-full bg-emerald-700 hover:bg-emerald-800 text-white font-bold py-3 px-6 rounded-lg transition-colors disabled:opacity-70 flex items-center justify-center gap-2"
            >
              {loading ? (
                <>
                  <svg className="animate-spin h-5 w-5 text-white" viewBox="0 0 24 24">
                    <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" fill="none"></circle>
                    <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Processing...</span>
                </>
              ) : (
                <span>Shorten URL</span>
              )}
            </button>
          </form>

          {error && (
            <div className="mt-6 p-4 bg-red-50 border border-red-200 rounded-lg">
              <p className="text-red-700 text-sm text-center">{error}</p>
            </div>
          )}

          {shortUrl && (
            <div className="mt-6 p-6 bg-emerald-50 rounded-xl border border-emerald-200">
              <div className="flex items-center gap-2 mb-4">
                <svg className="w-5 h-5 text-emerald-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
                </svg>
                <p className="font-bold text-emerald-900 text-sm">Success! Your Short Link:</p>
              </div>

              <div className="flex gap-2 mb-6">
                <input
                  type="text"
                  value={shortUrl}
                  readOnly
                  className="flex-1 px-3 py-2 bg-white border border-emerald-200 rounded text-sm text-emerald-800 font-mono focus:outline-none"
                />
                <button
                  onClick={copyToClipboard}
                  className="px-4 py-2 bg-emerald-700 hover:bg-emerald-800 text-white text-sm font-bold rounded transition-colors"
                >
                  {copied ? 'Copied' : 'Copy'}
                </button>
              </div>

              <div className="space-y-4 pt-4 border-t border-emerald-200">
                <div className="flex justify-between items-center">
                  <div>
                    <p className="text-[10px] font-bold text-emerald-800 uppercase">Total Clicks</p>
                    <p className="text-xl font-black text-emerald-900">{clicks}</p>
                  </div>
                  <button
                    onClick={refreshStats}
                    disabled={checkingStats}
                    className="text-xs font-bold text-emerald-700 hover:underline disabled:text-gray-400"
                  >
                    {checkingStats ? 'Updating...' : 'Update Stats'}
                  </button>
                </div>
                
                <div className="grid grid-cols-1 gap-2 text-xs">
                  <div>
                    <span className="font-bold text-emerald-800 uppercase block text-[10px]">Original</span>
                    <p className="text-gray-600 truncate">{originalUrl}</p>
                  </div>
                  <div>
                    <span className="font-bold text-emerald-800 uppercase block text-[10px]">Created At</span>
                    <p className="text-gray-600">{createdAt}</p>
                  </div>
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;