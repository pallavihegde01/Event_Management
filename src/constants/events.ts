import { server } from "./server";

export type EventStatus = "Upcoming" | "Postponed" | "Completed" | "Cancelled";

export type EventMonth =
  | "January"
  | "February"
  | "March"
  | "April"
  | "May"
  | "June"
  | "July"
  | "August"
  | "September"
  | "October"
  | "November"
  | "December";

export interface Event {
  id: number;
  title: string;
  description: string;
  start_time: string;
  end_time: string;
  status: EventStatus;
}

export async function getEvents(): Promise<Event[]> {
  try {
    const res = await fetch(`${server}/api/v1/events`);

    if (!res.ok) {
      throw new Error("Failed to fetch events");
    }

    const data = await res.json();

    return data;
  } catch (e) {
    throw "Failed to fetch events";
  }
}