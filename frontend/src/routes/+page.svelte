<script>
  import { onMount } from "svelte";
  let counter = 0;

  // Function to fetch the counter value from the server
  async function fetchCounter() {
    try {
      const response = await fetch("http://localhost:8080/value");
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      counter = data.value; // Set the counter value to the response
    } catch (error) {
      console.error('Error fetching counter:', error);
    }
  }

  // Function to increment the counter value on the server
  async function incrementCounter() {
    try {
        const response = await fetch("http://localhost:8080/increment", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            }
        });

        if (!response.ok) throw new Error('Network response was not ok');

        const data = await response.json();
        counter = data.value;
    } catch (error) {
        console.error('Error incrementing counter:', error);
    }
}

  // Fetch counter value when the component is mounted
  onMount(() => {
    fetchCounter();
    // Polling every 1 second to fetch the counter value
    setInterval(fetchCounter, 100);
  });
</script>

<style>
  /* Ensure full height for centering */
  html, body {
    height: 100%;
    margin: 0;
  }

  /* Centering container */
  .container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh; /* Full viewport height */
    flex-direction: column; /* Aligns the button and counter vertically */
    background-color: rgba(255, 255, 255, 0);
  }

  /* Button styling */
  .button {
    background: red;
    color: white;
    font-size: 24px;
    border: none;
    padding: 15px 30px;
    border-radius: 12px;
    cursor: pointer;
    box-shadow: 0 4px 10px rgba(255, 0, 0, 0.5);
    transition: transform 0.1s ease, box-shadow 0.1s ease;
  }

  /* Button hover effect */
  .button:hover {
    transform: scale(1.05);
    box-shadow: 0 6px 15px rgba(255, 0, 0, 0.7);
  }

  /* Counter styling */
  .counter {
    font-size: 48px;
    margin-bottom: 20px;
    color: black;
  }
</style>

<div class="container">
  <!-- Display counter value -->
  <div class="counter">{counter}</div>

  <!-- Button to increment the counter -->
  <button class="button" on:click={incrementCounter}>+1</button>
</div>
