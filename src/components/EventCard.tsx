import StatusBadge from "./StatusBadge";
import QASection from "./QASection";
import { EventType } from "../constants/eventPage";

interface EventCardProps {
  event: EventType;
}

function EventCard({ event }: EventCardProps) {
  return (
    <div className="w-full bg-white/80 backdrop-blur-xl shadow-2xl rounded-3xl overflow-hidden border border-blue-100">
      <div className="bg-gradient-to-r from-blue-600 to-indigo-700 text-white p-8">
        <div className="flex justify-between items-start flex-wrap gap-4">
          <div>
            <h1 className="text-3xl md:text-4xl font-bold leading-tight">
              {event.title}
            </h1>
            <p className="mt-2 text-blue-100 text-lg">
              Hosted by <span className="font-semibold">{event.company}</span>
            </p>
          </div>

          <StatusBadge status={event.status} />
        </div>

        {event.verified && (
          <p className="mt-3 text-sm font-semibold text-green-200">
            ✔ Verified Event
          </p>
        )}
      </div>
      <div className="p-8 space-y-8">
        <div>
          <h2 className="text-xl font-semibold mb-2 text-gray-800">
            About the Event
          </h2>
          <p className="text-gray-600 leading-relaxed">{event.description}</p>
        </div>
        <div className="grid sm:grid-cols-2 gap-6 text-sm">
          <div className="bg-blue-50 p-4 rounded-xl">
            <p className="font-semibold text-gray-700">Date</p>
            <p className="text-gray-600">{event.date}</p>
          </div>

          <div className="bg-blue-50 p-4 rounded-xl">
            <p className="font-semibold text-gray-700">Time</p>
            <p className="text-gray-600">{event.time}</p>
          </div>

          <div className="bg-blue-50 p-4 rounded-xl">
            <p className="font-semibold text-gray-700">Venue</p>
            <p className="text-gray-600">{event.venue}</p>
          </div>

          <div className="bg-blue-50 p-4 rounded-xl">
            <p className="font-semibold text-gray-700">Organizer</p>
            <p className="text-gray-600">{event.organizers}</p>
          </div>
        </div>
        <div className="text-center">
          <a
            href={event.bookingLink}
            target="_blank"
            rel="noopener noreferrer"
            className="inline-block bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-full text-base font-semibold transition duration-300 shadow-lg hover:shadow-xl"
          >
            Register Now
          </a>
        </div>
        <p className="text-xs text-gray-400 text-center">
          Source: {event.source}
        </p>
        <QASection faqs={event.faqs} />
      </div>
    </div>
  );
}

export default EventCard;
