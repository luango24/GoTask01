import http from 'k6/http';
import { sleep } from 'k6';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

export default function () {
    http.get('http://localhost:5000/');
  sleep(1);
}
export const options = {
    stages: [
        { duration: '5s', target: 3 }, // below normal load
        { duration: '5s', target: 4 }, // normal load
        { duration: '8s', target: 5 }, // around the breaking point
        { duration: '10s', target: 8 },
        { duration: '12s', target: 15 }, // beyond the breaking point
        { duration: '12s', target: 15 },
        { duration: '5s', target: 0 }, // scale down. Recovery stage.
    ],
};

export function handleSummary(data) {
    return {
        "summary.html": htmlReport(data),
    };
}