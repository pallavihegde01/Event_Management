
export type EventStatus =
    "Upcoming"
    | "Postponed"
    | "Completed"
    | "Cancelled";

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
    month: EventMonth;
    postedDate: string;
    updatedDate: string;
    status: EventStatus;
}

export const events: Event[] = [
    {
        id: 1,
        title: "React Meetup 2026",
        description: "Join us for an exciting React community meetup.",
        month: "February",
        postedDate: "Feb 10, 2026",
        updatedDate: "Feb 15, 2026",
        status: "Upcoming",
    },
    {
        id: 2,
        title: "Startup Networking Event",
        description: "Connect with founders and tech enthusiasts.",
        month: "March",
        postedDate: "Mar 05, 2026",
        updatedDate: "Mar 12, 2026",
        status: "Postponed",
    },
    {
        id: 3,
        title: "Tech Growth Summit",
        description: "Learn scaling strategies from industry leaders.",
        month: "May",
        postedDate: "May 05, 2026",
        updatedDate: "May 12, 2026",
        status: "Cancelled",
    },
    {
        id: 4,
        title: "New Year Hackathon",
        description: "Start your year with innovation and coding.",
        month: "January",
        postedDate: "Jan 05, 2026",
        updatedDate: "Jan 12, 2026",
        status: "Completed",
    },
];
