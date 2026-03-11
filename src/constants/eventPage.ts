import { server } from "./server";

export type EventStatus = "Ongoing" | "Completed" | "Postponed" | "Upcoming";

export interface FAQ {
  question: string;
  answer: string;
}

export interface EventType {
  id: number;
  title: string;
  company: string;
  description: string;
  date: string;
  time: string;
  status: EventStatus;
  venue: string;
  organizers: string;
  bookingLink: string;
  source: string;
  verified: boolean;
  faqs: FAQ[];
}

export async function getEvents(): Promise<EventType[]> {
  try {
    const res = await fetch(`${server}/api/v1/events`);
    if (!res.ok) throw new Error("Failed to fetch");
    const raw = await res.json();
    return raw.map((e: any) => ({
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
    }));
  } catch (e) {
    throw "Failed to fetch events";
  }
}

export async function createEvent(event: Omit<EventType, "id">): Promise<EventType> {
  const res = await fetch(`${server}/api/v1/events`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      ...event,
      start_time: new Date().toISOString(),
      end_time: new Date().toISOString(),
    }),
  });
  if (!res.ok) throw "Failed to create event";
  return res.json();
}