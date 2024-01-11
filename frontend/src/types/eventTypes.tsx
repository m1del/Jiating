export type Admin = {
  id: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
  name: string;
  email: string;
  position: string;
  status: string;
  events: Array<EventAuthor>;
};

export type EventData = {
  id: string;
  created_at: string | null;
  updated_at: string | null;
  event_name: string;
  date: string;
  description: string;
  content: string;
  is_draft: boolean;
  published_at: string | null;
  images: Array<EventImage>;
  authors: Array<Admin>;
};

export type EventAuthor = {
  admin_id: string;
  event_id: string;
};

export type EventImage = {
  id: string; // primary key, UUID
  created_at: string; // time of creation
  updated_at: string; // time of last update
  image_url: string; // url of the image (s3)
  is_display: boolean; // if the image is the display image for the event
};

export type UpdateEventRequest = {
  updated_data: object;
  new_images: Array<EventImage>;
  removed_image_ids: Array<string>;
  new_display_image_id: string;
  editor_admin_id: string;
};

export type CreateEventRequest = {
  event_name: string;
  date: string;
  description: string;
  content: string;
  is_draft: boolean;
  images: Array<EventImage>;
  author_ids: Array<string>;
};

