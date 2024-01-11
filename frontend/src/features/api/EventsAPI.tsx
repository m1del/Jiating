import {
  CreateEventRequest,
  UpdateEventRequest,
  EventData,
  EventImage,
} from '../../components/events/EventModel';

const getAdminByEmail = async (
  adminEmail: string | undefined,
): Promise<string> => {
  if (!adminEmail) {
    return Promise.reject(new Error('Admin email is undefined'));
  }
  try {
    const resp = await fetch(
      `http://localhost:3000/api/admin/get-by-email/${adminEmail}`,
      {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      },
    );
    if (!resp.ok) {
      throw new Error('Failed to get admin info');
    } else {
      const respData = await resp.json();
      return respData.id;
    }
  } catch (err) {
    return Promise.reject(err);
  }
};

const createEvent = async (
  eventData: EventData,
  adminID: string,
): Promise<string> => {
  try {
    const createData: CreateEventRequest = {
      event_name: eventData.event_name,
      date: eventData.date,
      description: eventData.description,
      content: eventData.content,
      is_draft: eventData.is_draft,
      images: eventData.images,
      author_ids: [adminID],
    };

    const resp = await fetch('http://localhost:3000/api/event/create', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify(createData),
    });
    if (!resp.ok) {
      throw new Error('Failed to send post');
    }
    const responseData = await resp.json();
    return responseData.id;
  } catch (err) {
    return Promise.reject(err);
  }
};

const updateEvent = async (
  eventData: EventData,
  adminID: string,
  oldImages: Array<EventImage>,
): Promise<Response> => {
  try {
    const updateData: UpdateEventRequest = {
      updated_data: {
        event_name: eventData.event_name,
        date: eventData.date,
        description: eventData.description,
        content: eventData.content,
        is_draft: eventData.is_draft,
      },
      new_images: eventData.images,
      removed_image_ids: oldImages.map((image) => image.id),
      new_display_image_id: '',
      editor_admin_id: adminID,
    };

    const resp = await fetch(
      `http://localhost:3000/api/event/update/${eventData.id}`,
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify(updateData),
      },
    );
    if (!resp.ok) {
      throw new Error('Failed to update post');
    }
    return resp;
  } catch (err) {
    return Promise.reject(err);
  }
};

//TODO fix when backend changes
// if (imageRef.current && imageRef.current.files) {
//   try {
//     const resp2 = await fetch(
//       `http://localhost:3000/api/event/upload/${eventID}/${imageRef.current.files[0]}`,
//       {
//         method: 'POST',
//         headers: { 'Content-Type': 'application/json' },
//         credentials: 'include',
//       },
//     );
//     if (!resp2.ok) {
//       throw new Error('Failed to get presigned url for image');
//     } else {
//       const resp2data = await resp2.json();
//       try {
//         const resp3 = await fetch(resp2data.url, {
//           method: 'PUT',
//           headers: { 'Content-Type': 'application/xml' },
//           body: imageRef.current.files[0],
//         });
//         if (!resp3.ok) {
//           throw new Error('Failed to upload image');
//         }
//       } catch (err) {
//         console.log(err);
//       }
//       console.log(resp2.json());
//     }
//   } catch (err) {
//     console.log(err);
//   }
// }

export { getAdminByEmail, createEvent, updateEvent };
