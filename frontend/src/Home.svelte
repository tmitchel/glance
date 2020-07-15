<script>
  let hey = "hey";

  let cardTitle = "";
  let cardContent = "";
  let cardCreator = "11111111-1111-1111-1111-111111111111";

  import { CreateService, GetService } from "../../generated/oto.gen.js";

  // create services
  const createService = new CreateService();
  const getService = new GetService();

  let cardPromise = async function() {
    try {
      return await getService.cards({});
    } catch(err) {
      console.log(err)
    }
  }()

  let userPromise = async function() {
    try {
      return await getService.users({});
    } catch(err) {
      console.log(err)
    }
  }()

  // createService.card({
  //   title: "card title",
  //   content: "card content",
  //   Creator: "11111111-1111-1111-1111-111111111111"
  // });
  
  function color(status) {
    if (status === 1) {
      return "green";
    } else if (status === 2) {
      return "red";
    } else if (status === 3) {
      return "blue";
    } else if (status === 4) {
      return "black";
    }
    return "white";
  }

</script>

<div>
  {#await userPromise}
    <div class="ui segment">
      <div class="ui active dimmer">
        <div class="ui text loader">Loading</div>
      </div>
    </div>
  {:then msg}
    {#each msg.users as user}
      <div class="ui segment">
        <!-- <h1>{user.name}</h1> -->
        <h1>User Name</h1>
      </div>
    {/each}
  {:catch error}
    <h1>{error}</h1>
  {/await}
</div>

<div class="ui container">
  <form class="ui form">
    <div class="field">
      <label>Title</label>
      <input type="text" bind:value={cardTitle} placeholder="Task Title">
    </div>
    <div class="field">
      <label>Description</label>
      <textarea bind:value={cardContent} rows=4 placeholder="Describe the task..."></textarea>
    </div>
    <button class="ui button" type="submit">Submit</button>
  </form>
</div>

<div class="ui cards">
  {#await cardPromise}
    <div class="ui segment">
      <div class="ui active dimmer">
        <div class="ui text loader">Loading</div>
      </div>
    </div>
  {:then msg}
    {#each msg.cards as card}
      <div class="{color(card.status)} card">
        <div class="custom content">
          <h1 class="header">{card.title}</h1>
          <div class="custom-text description">
            {card.content}
          </div>
        </div>
      </div>
    {/each}
  {:catch error}
    <h1>{error}</h1>
  {/await}
</div>