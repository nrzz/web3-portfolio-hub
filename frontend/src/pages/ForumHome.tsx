import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

export const ForumHome: React.FC = () => {
  const [questions, setQuestions] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchQuestions = async () => {
      setLoading(true);
      setError(null);
      try {
        const res = await fetch('/api/v1/forum/questions', {
          headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
        });
        if (!res.ok) throw new Error('Failed to fetch questions');
        const data = await res.json();
        setQuestions(data.questions || []);
      } catch (err: any) {
        setError(err.message || 'Unknown error');
      } finally {
        setLoading(false);
      }
    };
    fetchQuestions();
  }, []);

  return (
    <div className="max-w-3xl mx-auto py-8">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Community Forum</h1>
        <Link to="/forum/ask" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition">Ask Question</Link>
      </div>
      {loading ? (
        <div className="text-gray-500">Loading questions...</div>
      ) : error ? (
        <div className="text-red-600">{error}</div>
      ) : (
        <div className="bg-white shadow rounded-lg divide-y">
          {questions.length === 0 ? (
            <div className="px-6 py-4 text-gray-500">No questions yet.</div>
          ) : questions.map((q: any) => (
            <Link to={`/forum/${q.id}`} key={q.id} className="block px-6 py-4 hover:bg-gray-50 transition">
              <div className="font-semibold text-lg">{q.title}</div>
              <div className="text-sm text-gray-500">By {q.author || q.user_id} on {q.createdAt || q.created_at}</div>
            </Link>
          ))}
        </div>
      )}
    </div>
  );
}; 