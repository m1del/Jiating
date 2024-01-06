import React, { FormEvent, useState, useRef } from 'react';

export default function PostForm() {
  const submitRef = useRef<HTMLFormElement>(null);

  type FormData = {
    author: string;
    event: string;
    description: string;
    date: string;
    draft: boolean;
  };

  const [formData, setFormData] = useState<FormData>({
    author: '',
    event: '',
    description: '',
    date: '',
    draft: false,
  });

  const handleSubmit = async () => {
    // text Fields Validation
    if (
      !sanitizeInput(formData.author) ||
      !sanitizeInput(formData.event) ||
      !sanitizeInput(formData.description)
    ) {
      alert('Invalid characters in the input fields.');
      return;
    }

    // length Checks
    if (
      formData.author.length > 50 ||
      formData.event.length > 50 ||
      formData.description.length > 2000
    ) {
      alert('Input is too long.');
      return;
    }

    console.log(formData);

    try {
      const resp = await fetch('http://localhost:3000/admin/publish-form', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify(formData),
      });
      if (!resp.ok) {
        throw new Error('Failed to send post');
      }
      console.log(resp);
      //TODO: handle sucess -> clear form & show sucess message
    } catch (err) {
      console.error(err);
    }
  };

  const checkDraft = (e: FormEvent, draft: boolean) => {
    e.preventDefault();
    setFormData({ ...formData, draft: draft });
    if (submitRef.current) {
      submitRef.current.checkValidity();
      handleSubmit();
    }
  };

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const sanitizeInput = (input: string) => {
    const re = /[<>]/;
    return !re.test(input);
  };

  // Needs to have EventName, EventDate, EventContent (embed links), Draft bool, Admin who wrote, Images
  return (
    <div className="mt-16 flex h-screen w-full flex-col items-center bg-cyan-700">
      <div className="flex flex-col items-center rounded-lg  bg-white px-36 pb-24 pt-10">
        <h2 className="mb-10 text-2xl font-semibold">Create Post</h2>
        <form
          ref={submitRef}
          onSubmit={handleSubmit}
          className="flex flex-col gap-2"
        >
          <h3 className="text-sm font-semibold">Author and Event Name *</h3>
          <div className="flex gap-7">
            <input
              required
              type="text"
              name="author"
              value={formData.author}
              onChange={handleChange}
              placeholder="Author Name"
              className="border-2 border-gray-300 px-2 py-2"
            />
            <input
              required
              type="text"
              name="event"
              value={formData.event}
              onChange={handleChange}
              placeholder="Event Name"
              className="border-2 border-gray-300 px-2 py-2"
            />
          </div>
          <h3 className="text-sm font-semibold">Article Description *</h3>
          <textarea
            required
            name="description"
            placeholder="Description"
            value={formData.description}
            onChange={handleChange}
            className="border-2 border-gray-300 px-2 pt-2"
          ></textarea>
          <span>
            <input
              type="date"
              name="date"
              value={formData.date}
              onChange={handleChange}
            />
          </span>
          <input type="file" name="image" accept="image/png, image/jpeg" />

          <span className="mx-auto flex gap-3">
            <input
              onClick={(e) => checkDraft(e, true)}
              type="submit"
              value="Preview"
              className="rounded-lg bg-gray-400 px-3 py-1 text-white"
            />
            <input
              onClick={(e) => checkDraft(e, false)}
              type="submit"
              value="Publish"
              className="rounded-lg bg-gray-400 px-3 py-1 text-white"
            />
          </span>
        </form>
      </div>
    </div>
  );
}
