<script>
  let hey = "hey";

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

  // createService.card({
  //   title: "card title",
  //   content: "card content",
  //   Creator: "11111111-1111-1111-1111-111111111111"
  // });
  
  // const cardService = new CardService();
  // let meh = {};
  // cardService
  //   .cards({})
  //   .then(resp => {
  //     meh = resp;
  //   })
  //   .catch(err => console.log(err));

  //   cardService.new({
  //       title: "MEH",
  //       Content: "hey",
  //       Creator: "a"
  //   }).then(resp => console.log(resp)).catch(err => console.log(err))
</script>

<div class="ui cards">
  {#await cardPromise}
    <div class="ui segment">
      <div class="ui active dimmer">
        <div class="ui text loader">Loading</div>
      </div>
    </div>
  {:then msg}
    {#each msg.cards as card}
      <div class="card">
        <div class="content">
          <h1 class="header">{card.title}</h1>
          <div class="description">
            {card.content}
          </div>
        </div>
      </div>
    {/each}
  {:catch error}
    <h1>{error}</h1>
  {/await}
</div>

<h1>{hey}</h1>
<h2>mey</h2>
<!-- <pre>{meh}</pre> -->
<h1>you</h1>
