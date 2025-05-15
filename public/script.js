async function shortenUrl() {
  const longUrl = document.getElementById("longUrl").value;
  const response = await fetch("/shorten", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ url: longUrl }),
  });

  const result = await response.json();
  document.getElementById("result").innerText = `Short URL: ${result.short_url}`;
  document.getElementById("copyBtn").style.display = "inline-block";
}

function copyToClipboard() {
  const text = document.getElementById("result").innerText.replace("Short URL: ", "");
  navigator.clipboard.writeText(text).then(() => {
    alert("Short URL copied to clipboard!");
  }).catch(err => {
    alert("Failed to copy!");
  });
}
