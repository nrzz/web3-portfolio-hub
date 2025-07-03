import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

export const QuestionDetail: React.FC = () => {
  const { id } = useParams();
  const [question, setQuestion] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchQuestion = async () => {
      setLoading(true);
      setError(null);
      try {
        const res = await fetch(`/api/v1/forum/questions/${id}`, {
          headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
        });
        if (!res.ok) throw new Error('Failed to fetch question');
        const data = await res.json();
        setQuestion(data.question);
      } catch (err: any) {
        setError(err.message || 'Unknown error');
      } finally {
        setLoading(false);
      }
    };
    fetchQuestion();
  }, [id]);

  return (
    <div className="max-w-3xl mx-auto py-8">
      {loading ? (
        <div className="text-gray-500">Loading question...</div>
      ) : error ? (
        <div className="text-red-600">{error}</div>
      ) : question ? (
        <div className="bg-white shadow rounded-lg p-6">
          <h1 className="text-2xl font-bold mb-2">{question.title}</h1>
          <div className="text-gray-700 mb-4 whitespace-pre-line">{question.body}</div>
          <div className="text-sm text-gray-500 mb-2">By {question.author || question.user_id} on {question.createdAt || question.created_at}</div>
          <div className="mb-4">
            {question.tags && question.tags.length > 0 && (
              <div className="flex flex-wrap gap-2">
                {question.tags.map((tag: any) => (
                  <span key={tag.id || tag} className="bg-blue-100 text-blue-800 px-2 py-1 rounded text-xs">{tag.name || tag}</span>
                ))}
              </div>
            )}
          </div>
          {/* Placeholder for answers and comments */}
          <div className="mt-8">
            <h2 className="text-xl font-semibold mb-2">Answers</h2>
            <div className="text-gray-500">(Answers will appear here)</div>
          </div>
          <div className="mt-6">
            <h2 className="text-lg font-semibold mb-2">Add an Answer</h2>
            <div className="text-gray-500">(Answer form will go here)</div>
          </div>
        </div>
      ) : null}
    </div>
  );
}; 