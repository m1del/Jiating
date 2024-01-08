type EventData = {
  id: string;
  created_at: string | null;
  updated_at: string | null;
  admin_id: number;
  event_name: string;
  date: string;
  description: string;
  content: string;
  is_draft: boolean;
  published_at: string | null;
  image_id: number;
};

export type { EventData };
