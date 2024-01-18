import React, { useState } from 'react';
import Editor from './form/Editor';

// TODO: define the types in src/types/
type EventImage = {};
type Admin = {};

type EventFormProps = {
  onSubmit: (event: Event) => void;
};

type Event = {
  id: string;
  createdAt: Date;
  updatedAt: Date;
  eventTitle: string;
  metaTitle: string;
  slug: string;
  date: Date;
  description: string;
  content: string;
  isDraft: boolean;
  publishedAt: Date | null;
  images: EventImage[];
  authors: Admin[];
};

const EventForm: React.FC<EventFormProps> = ({ onSubmit }) => {
  const [event, setEvent] = useState<Event>({
    // initialize state with default values or fetch existing data if updating?
  });


  const [editorContent, setEditorContent] = useState<string>('');

  const handleEditorContentChange = (content: string) => {
    setEditorContent(content);
  };

  const handleSubmit = (event: React.FormEvent) => { // TODO: implement
    event.preventDefault();
    // use editorContent to submit to backend
    const content = editorContent;
    // parse content, find img with data.isdisplay set to true
    // extract the URL of the display image and other necessary data
    // Submit the data to  backend

    //TODO: figure out how to handle slice of images?
    console.log(content)
  }


  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setEvent({ ...event, [e.target.name]: e.target.value });
  };

//   const handleSubmit = (event: React.FormEvent) => {
//   event.preventDefault();
//   const content = quill?.root.innerHTML;
//   // parse content to find the image with data.IsDisplay set to true
//   // extract the URL of the display image and other necessary data
//   // Submit the data to  backend
// };

  return (
        <form onSubmit={handleSubmit} className="container mx-auto flex flex-col p-6 bg-white shadow-md rounded-lg">
            <label className="mb-4 font-semibold">
                Event Title:
                <input type="text" name="eventTitle" value={event.eventTitle} onChange={handleChange} className="w-full p-3 mt-2 border-2 border-gray-200 rounded-lg focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
            </label>
            <label className="mb-4 font-semibold">
                Meta Title:
                <input type="text" name="metaTitle" value={event.metaTitle} onChange={handleChange} className="w-full p-3 mt-2 border-2 border-gray-200 rounded-lg focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
            </label>
            <label className="mb-4 font-semibold">
                Slug:
                <input type="text" name="slug" value={event.slug} onChange={handleChange} className="w-full p-3 mt-2 border-2 border-gray-200 rounded-lg focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
            </label>
            <label className="mb-4 font-semibold">
                Description:
                <textarea name="description" value={event.description} onChange={handleChange} className="w-full p-3 mt-2 border-2 border-gray-200 rounded-lg h-32 focus:border-blue-500 focus:ring-1 focus:ring-blue-500" />
            </label>
            <label className="mb-4 font-semibold flex items-center">
                Is Draft:
                <input type="checkbox" name="isDraft" checked={event.isDraft} onChange={e => setEvent({ ...event, isDraft: e.target.checked })} className="ml-2 h-5 w-5" />
            </label>
            <div className="mx-auto mb-4">
                <Editor onContentChange={handleEditorContentChange}/>
            </div>
            <button type="submit" className="w-full p-3 bg-blue-500 hover:bg-blue-600 text-white rounded-lg transition-colors">Submit</button>
        </form>
    );
};

export default EventForm;
