import EventCard from "../components/EventCard";
import { eventPage } from "../constants/eventPage";
const EventPage = () => {
  const eventDetail = eventPage[0];

  return (
    <div className="min-h-screen bg-blue-50">
      <div className="bg-gradient-to-r from-blue-600 to-indigo-700 py-12 px-6 text-center text-white">
        <h1 className="text-4xl md:text-5xl font-bold">Event Details</h1>

        <p className="mt-3 text-blue-100 text-lg">
          Discover everything about this event
        </p>
      </div>

      <div className="max-w-5xl mx-auto px-6 py-12">
        <EventCard event={eventDetail} />
      </div>
    </div>
  );
};

export default EventPage;
