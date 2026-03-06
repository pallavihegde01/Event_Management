import { FAQ } from "../constants/eventPage";

interface QASectionProps {
  faqs: FAQ[];
}

function QASection({ faqs }: QASectionProps) {
  return (
    <div className="mt-4 space-y-3">
      <h3 className="font-semibold text-lg">Q & A</h3>

      {faqs.map((faq, index) => (
        <div key={index} className="bg-gray-50 p-3 rounded-lg">
          <p className="font-medium">Q: {faq.question}</p>
          <p className="text-gray-600">A: {faq.answer}</p>
        </div>
      ))}
    </div>
  );
}

export default QASection;
