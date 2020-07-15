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

  // createService.user({
  //   name: "Tyler Mitchell",
  //   email: "tylerbmitchell6@gmail.com",
  //   password: "test-password"
  // });

  // createService.card({
  //   title: "Test Card",
  //   content: "Test card content is here....",
  //   Creator: "888aabf9-6289-4b74-afc6-976a4a74ad4a"
  // });

  let userPromise = async function() {
    try {
      return await getService.homePage({email: "tylerbmitchell6@gmail.com"});
    } catch(err) {
      console.log(err)
    }
  }()
  
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

<h1>Users</h1>
<div class="ui five cards">
  {#await userPromise}
    <div class="ui segment">
      <div class="ui active dimmer">
        <div class="ui text loader">Loading</div>
      </div>
    </div>
  {:then msg}
    {#each msg.pairs as pair}
      <div class="ui card">
        <div class="content">
          <h1 class="header">{pair.User.name}</h1>
          <div class="ui {color(pair.Card.status)} card">
            <div class="content">
              <h3 class="header">{pair.Card.title || "card title"}</h3>
              <div class="description">
                {pair.Card.content}
              </div>
            </div>
          </div>
        </div>
      </div>
    {/each}
  {:catch error}
    <h1>{error}</h1>
  {/await}
</div>

<!-- <div class="ui container">
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
</div> -->

<h1>Unclaimed Cards</h1>
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