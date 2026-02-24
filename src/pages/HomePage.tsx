import React, { useState } from "react";
import {
  FaBell,
  FaSearch,
  FaCompass,
  FaUser,
  FaHeart,
  FaBookmark,
} from "react-icons/fa";
import { events } from "../constants/events";
import months from "../constants/months";

export default function HomePage() {
  const currentMonth = new Date().getMonth();
  const [selectedMonth, setSelectedMonth] = useState(months[currentMonth]);

  const filteredEvents = events.filter(
    (event) => event.month === selectedMonth,
  );

  const getStatusStyle = (status) => {
    switch (status.toLowerCase()) {
      case "upcoming":
        return "bg-green-100 text-green-700";
      case "postponed":
        return "bg-yellow-100 text-yellow-700";
      case "completed":
        return "bg-gray-200 text-gray-700";
      case "cancelled":
        return "bg-red-100 text-red-700";
      default:
        return "bg-gray-100 text-gray-600";
    }
  };

  return (
    <div className="flex flex-col h-screen bg-gray-100">
      {/* Navbar */}
      <nav className="flex justify-between items-center px-5 py-4 bg-gray-900 text-white shadow-md">
        <h2 className="text-xl font-semibold">Logo</h2>
        <FaBell className="text-lg cursor-pointer" />
      </nav>

      {/*Contents*/}
      <div className="flex flex-1 overflow-hidden">
        {/* Events Area */}
        <div className="flex-1 overflow-y-auto p-4 space-y-5">
          <h2 className="text-2xl font-bold mb-4">{selectedMonth}</h2>

          {filteredEvents.length === 0 ? (
            <p className="text-gray-500">No events for this month.</p>
          ) : (
            filteredEvents.map((event) => (
              <div key={event.id} className="bg-white rounded-xl shadow-sm p-4">
                {/* Section 1 */}
                <div className="flex justify-between items-center text-sm">
                  <span className="text-gray-500">{event.postedDate}</span>

                  <div className="flex items-center space-x-3">
                    <FaHeart className="cursor-pointer text-gray-500 hover:text-red-500" />
                    <FaBookmark className="cursor-pointer text-gray-500 hover:text-blue-500" />
                    <span
                      className={`px-2 py-1 rounded text-xs font-medium ${getStatusStyle(
                        event.status,
                      )}`}
                    >
                      {event.status}
                    </span>
                  </div>
                </div>

                {/* Section 2 */}
                <div className="flex mt-4 space-x-4">
                  <img
                    src="/event-image.png"
                    alt="event"
                    className="w-28 h-20 rounded-lg"
                  />
                  <div>
                    <h3 className="font-semibold text-lg">{event.title}</h3>
                    <p className="text-gray-600 text-sm mt-1">
                      {event.description}
                    </p>
                  </div>
                </div>

                {/* Section 3 */}
                <div className="flex justify-between items-center mt-4 text-sm">
                  <button className="bg-blue-600 text-white px-3 py-1 rounded-md hover:bg-blue-700">
                    More Details
                  </button>
                  <span className="text-gray-500">
                    Last Updated: {event.updatedDate}
                  </span>
                </div>
              </div>
            ))
          )}
        </div>

        <div className="fixed right-0 top-1/2 -translate-y-1/2 group flex items-center">
          <div className="w-1 h-40 bg-blue-600 rounded-l-full"></div>

          <div
            className="
      bg-white shadow-xl rounded-l-xl
      w-0 group-hover:w-32
      overflow-hidden
      transition-all duration-300
      h-64
  "
          >
            <div className="h-full overflow-y-auto p-3 space-y-2">
              {months.map((month) => (
                <button
                  key={month}
                  onClick={() => setSelectedMonth(month)}
                  className={`block w-full text-left text-sm px-2 py-1 rounded 
            ${
              selectedMonth === month
                ? "bg-blue-100 text-blue-600 font-semibold"
                : "hover:bg-gray-100"
            }`}
                >
                  {month}
                </button>
              ))}
            </div>
          </div>
        </div>
      </div>

      {/* Bottom Navbar */}
      <nav className="flex justify-around items-center py-3 bg-white border-t shadow-inner">
        <FaCompass className="text-gray-600 text-xl cursor-pointer" />
        <FaSearch className="text-blue-600 text-2xl cursor-pointer" />
        <FaUser className="text-gray-600 text-xl cursor-pointer" />
      </nav>
    </div>
  );
}
