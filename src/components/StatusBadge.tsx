import { EventStatus } from "../constants/events";

interface StatusBadgeProps {
  status: EventStatus;
}

function StatusBadge({ status }: StatusBadgeProps) {
  const colors: Record<EventStatus, string> = {
    Ongoing: "bg-green-100 text-green-700",
    Completed: "bg-gray-200 text-gray-700",
    Postponed: "bg-yellow-100 text-yellow-700",
    Upcoming: "bg-blue-100 text-blue-700",
  };

  return (
    <span className={`px-3 py-1 rounded-full text-xs font-semibold ${colors[status]}`}>
      {status}
    </span>
  );
}

export default StatusBadge;