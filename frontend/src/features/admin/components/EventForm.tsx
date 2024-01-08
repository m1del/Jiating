import React, {
  FormEvent,
  useState,
  useRef,
  useEffect,
  useCallback,
} from 'react';
import CheckAuth from '../CheckAuth';
import { useAuth } from '../../../context/AuthContext';
import { v4 as uuidv4 } from 'uuid';
import { EventData } from '../../../components/events/EventModel';
import GetEvents from '../../../components/events/GetEvents';

export default function PostForm() {
  const { authUser, setAuthUser, setIsLoggedin } = useAuth();

  const [formData, setFormData] = useState<EventData>({
    id: '',
    created_at: null,
    updated_at: null,
    admin_id: 1,
    event_name: '',
    date: '',
    description: '',
    content: '',
    is_draft: true,
    published_at: null,
    image_id: 0,
  });

  const [formUpdate, setFormUpdate] = useState(false);

  //Check auth and get google auth information

  const submitRef = useRef<HTMLFormElement>(null);

  const handleSubmit = useCallback(async () => {
    // text Fields Validation
    if (
      !sanitizeInput(formData.event_name) ||
      !sanitizeInput(formData.description) ||
      !sanitizeInput(formData.content)
    ) {
      alert('Invalid characters in the input fields.');
      return;
    }

    // length Checks
    if (formData.event_name.length > 50 || formData.description.length > 2000) {
      alert('Input is too long.');
      return;
    }

    console.log(formData);

    try {
      const resp = await fetch('http://localhost:3000/admin/create-event', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify(formData),
      });
      if (!resp.ok) {
        throw new Error('Failed to send post');
      }
      console.log(resp);
      window.location.href = `/admin/get-event?id=${formData.id}`;
    } catch (err) {
      console.error(err);
    }
  }, [formData]);

  useEffect(() => {
    CheckAuth(setAuthUser, setIsLoggedin);
    GetEvents(setFormData);
    if (formUpdate) {
      handleSubmit();
      setFormUpdate(false);
    }
  }, [setAuthUser, setIsLoggedin, formUpdate, handleSubmit]);

  const checkDraft = async (e: FormEvent, isDraft: boolean) => {
    e.preventDefault();

    // Set eventID if it is a new event
    if (formData.id === '') setFormData({ ...formData, id: uuidv4() });

    setFormData({ ...formData, is_draft: isDraft });

    if (submitRef.current) {
      submitRef.current.checkValidity();
      setFormUpdate(true);
      // handleSubmit();
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
    <div className=" flex w-full flex-col items-center bg-cyan-700">
      <div className="my-16 flex w-[75%] flex-col items-center  rounded-lg bg-white px-36 pb-24 pt-10">
        <h2 className="mb-10 text-2xl font-semibold">Create Event</h2>
        <form
          ref={submitRef}
          onSubmit={handleSubmit}
          className="flex w-full flex-col gap-2"
        >
          <h3 className="text-sm font-semibold">Event Name *</h3>
          <div className="flex gap-7">
            <input
              required
              type="text"
              name="event_name"
              value={formData.event_name}
              onChange={handleChange}
              placeholder="Event Name"
              className="border-2 border-gray-300 px-2 py-2"
            />
          </div>
          <h3 className="w-[80%] text-sm font-semibold">
            Article Description *
          </h3>
          <textarea
            required
            rows={3}
            name="description"
            placeholder="Description"
            value={formData.description}
            onChange={handleChange}
            className="border-2 border-gray-300 px-2 pt-2"
          ></textarea>
          <h3 className="w-[80%] text-sm font-semibold">Article Content *</h3>
          <textarea
            required
            rows={15}
            name="content"
            placeholder="Content"
            value={formData.content}
            onChange={handleChange}
            className="border-2 border-gray-300 px-2 pt-2"
          ></textarea>
          <h3 className="text-sm font-semibold">Event Date *</h3>
          <span className="mb-3">
            <input
              type="date"
              name="date"
              value={formData.date}
              onChange={handleChange}
            />
          </span>
          <input
            type="file"
            name="image"
            accept="image/png, image/jpeg"
            className="mb-5"
          />

          <span className=" flex gap-3">
            <input
              onClick={(e) => checkDraft(e, true)}
              type="submit"
              value="Save & Preview"
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
