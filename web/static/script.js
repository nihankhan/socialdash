const refreshButton = document.getElementById("refresh");

refreshButton.addEventListener("click", () => {
    fetch("127.0.0.1:8080") // Assuming your server endpoint is at /api/data
        .then(response => response.json())
        .then(data => {
            document.querySelector(".platforms p:nth-child(1)").textContent = `Followers: ${data.FacebookCount}`;
            document.querySelector(".platforms p:nth-child(3)").textContent = `Followers: ${data.TwitterCount}`;
            document.querySelector(".platforms p:nth-child(5)").textContent = `Subscribers: ${data.YouTubeSubscribers}`;
            document.querySelector(".platforms p:nth-child(7)").textContent = `Followers: ${data.InstagramFollowers}`;
        })
        .catch(error => {
            console.error("Error fetching data:", error);
            // Display an error message to the user
        });
});
