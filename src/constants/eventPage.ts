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
    var data: EventType[] = [];
    try {
        const res = await fetch(`${server}/api/v1/events`);

    } catch (e) {
        throw "Failed to fetch events";
    }
    return data;
}

export const eventPage: EventType[] = [
    {
        id: 1,
        title: "AI & Future Tech Summit",
        company: "Code Technologies",
        description:
            "A full-day summit exploring AI, Web3, and Cloud innovation. Includes live demos, networking sessions, and internship opportunities.",
        date: "25 March 2026",
        time: "10:00 AM – 4:00 PM",
        status: "Ongoing",
        venue: "Microsoft Reactor, Richmond Circle, Bangalore",
        organizers: "EventPro Solutions",
        bookingLink: "https://example.com/book-ai-summit",
        source: "Tech Events Weekly Newsletter",
        verified: true,
        faqs: [
            {
                question: "Is this event free?",
                answer: "Yes, but registration is mandatory.",
            },
            {
                question: "Will certificates be provided?",
                answer: "Yes, participation certificates will be given.",
            },
        ],
    },
];
