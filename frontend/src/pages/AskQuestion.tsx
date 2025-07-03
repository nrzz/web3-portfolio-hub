import React, { useState } from 'react';

export const AskQuestion: React.FC = () => {
  const [title, setTitle] = useState('');
  const [body, setBody] = useState('');
  const [tags, setTags] = useState('');
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSuccess(false);
    try {
      const res = await fetch('/api/v1/forum/questions', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
        body: JSON.stringify({
          title,
          body,
          tags: tags.split(',').map(t => t.trim()).filter(Boolean)
        })
      });
      if (!res.ok) throw new Error('Failed to post question');
      setSuccess(true);
      setTitle('');
      setBody('');
      setTags('');
    } catch (err: any) {
      setError(err.message || 'Unknown error');
    }
  };

  return (
    <div className="max-w-xl mx-auto py-8">
      <h1 className="text-2xl font-bold mb-6">Ask a Question</h1>
      {success && (
        <div className="bg-green-100 text-green-800 p-4 rounded mb-4">Your question has been posted!</div>
      )}
      {error && (
        <div className="bg-red-100 text-red-800 p-4 rounded mb-4">{error}</div>
      )}
      <form onSubmit={handleSubmit} className="space-y-4 bg-white shadow rounded-lg p-6">
        <div>
          <label className="block font-medium mb-1">Title</label>
          <input type="text" className="w-full border rounded px-3 py-2" value={title} onChange={e => setTitle(e.target.value)} required />
        </div>
        <div>
          <label className="block font-medium mb-1">Body</label>
          <textarea className="w-full border rounded px-3 py-2" rows={6} value={body} onChange={e => setBody(e.target.value)} required />
        </div>
        <div>
          <label className="block font-medium mb-1">Tags (comma separated)</label>
          <input type="text" className="w-full border rounded px-3 py-2" value={tags} onChange={e => setTags(e.target.value)} />
        </div>
        <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition">Post Question</button>
      </form>
    </div>
  );
}; 