import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import EventCard from "../components/EventCard";
import { EventType } from "../constants/eventPage";
import { server } from "../constants/server";

const EventPage = () => {
  const { id } = useParams();
  const [event, setEvent] = useState<EventType | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    fetch(`${server}/api/v1/events/${id}`)
      .then((res) => {
        if (!res.ok) throw new Error();
        return res.json();
      })
      .then((e) => setEvent({
        id: e.id,
        title: e.title,
        company: e.company ?? "",
        description: e.description ?? "",
        date: new Date(e.start_time).toLocaleDateString("en-IN", {
          day: "numeric", month: "long", year: "numeric",
        }),
        time: `${new Date(e.start_time).toLocaleTimeString()} – ${new Date(e.end_time).toLocaleTimeString()}`,
        status: e.status,
        venue: e.venue ?? "",
        organizers: e.organizers ?? "",
        bookingLink: e.booking_link ?? "",
        source: e.source ?? "",
        verified: e.verified ?? false,
        faqs: e.faqs ?? [],
      }))
      .catch(() => setError("Failed to load event"))
      .finally(() => setLoading(false));
  }, [id]);

  if (loading)
    return (
      <div className="min-h-screen bg-blue-50 flex items-center justify-center">
        <p className="text-blue-600 text-lg font-medium">Loading event...</p>
      </div>
    );

  if (error || !event)
    return (
      <div className="min-h-screen bg-blue-50 flex items-center justify-center">
        <p className="text-red-500 text-lg">{error || "Event not found"}</p>
      </div>
    );

  return (
    <div className="min-h-screen bg-blue-50">
      <div className="bg-gradient-to-r from-blue-600 to-indigo-700 py-12 px-6 text-center text-white">
        <h1 className="text-4xl md:text-5xl font-bold">Event Details</h1>
        <p className="mt-3 text-blue-100 text-lg">
          Discover everything about this event
        </p>
      </div>
      <div className="max-w-5xl mx-auto px-6 py-12">
        <EventCard event={event} />
      </div>
    </div>
  );
};

export default EventPage;
