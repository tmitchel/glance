<script>
  import Task from "./Task.svelte";
  import Create from "./Create.svelte";
  import Summary from "./Summary.svelte";
  import Unclaimed from "./Unclaimed.svelte";
  import { CreateService, GetService } from "../../generated/oto.gen.js";

  const socket = new WebSocket("ws://localhost:8080/ws");
  socket.onmessage = (msg) => {
    console.log(msg.data);
    if (msg.data === "new-card") {
      cardPromise = getCard();
    } else if (msg.data === "new-user") {

    } else if (msg.data === "claim-card") {
      cardPromise = getCard();
    } else if (msg.data === "update-status") {
      userPromise = getUser();
    }
  };

  // create services
  export let createService = new CreateService();
  export let getService = new GetService();
  export let card = {
    title: "",
    content: "",
    status: 0
  };

  async function getCard() {
    try {
      return await getService.cards({});
    } catch(err) {
      console.log(err)
    }
  }

  $: cardPromise = getCard()

  async function getUser() {
    try {
      return await getService.homePage({email: "tylerbmitchell6@gmail.com"});
    } catch(err) {
      console.log(err);
    }
  }

  let userPromise = getUser();
  
  function color(status) {
    if (status === 1) {
      return "green-bkg";
    } else if (status === 2) {
      return "red-bkg";
    } else if (status === 3) {
      return "blue-bkg";
    } else if (status === 4) {
      return "black";
    }
    return "";
  }

  function cardColumns(n) {
    if (n < 4) {
      return "";
    } else if (n === 4) {
      return "four";
    }
    return "five";
  }

  function openTask(task) {
    card = task;
    window.$('#task.ui.modal').modal('show');
  }

  let volunteer = async function(userID, cardID) {
    try {
      createService.claimCard({userID: "888aabf9-6289-4b74-afc6-976a4a74ad4a", cardID: cardID})
    } catch(err) {
      console.log(err);
    }
  }

</script>

<!-- Bring everything in -->
<svelte:component this={Summary} cardColumns={cardColumns} userPromise={userPromise} color={color} openTask={openTask}></svelte:component>
<svelte:component this={Unclaimed} cardColumns={cardColumns} cardPromise={cardPromise} openTask={openTask}></svelte:component>
<svelte:component this={Create} createService={createService}></svelte:component>
<svelte:component this={Task} card={card} volunteer={volunteer}  createService={createService} usersCardID={userPromise}></svelte:component>
