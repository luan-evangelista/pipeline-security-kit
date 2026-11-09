declare const logger: any;
const payment = { cardNumber: "", cvv: "", id: "" };
const cardNumber = payment.cardNumber;

// ruleid: pan-in-logs
console.log("charging card", cardNumber);
// ruleid: pan-in-logs
console.error("declined", payment.cvv);
// ruleid: pan-in-logs
logger.info({ msg: "auth" }, payment.cardNumber);

// ok: pan-in-logs
console.log("charging payment", payment.id);
// ok: pan-in-logs
logger.info("masked card", maskPan(cardNumber));

function maskPan(pan: string) {
  return pan.slice(-4).padStart(pan.length, "*");
}
