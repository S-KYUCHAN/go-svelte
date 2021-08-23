<script>
    import { useNavigate, useLocation } from "svelte-navigator";
    import axios from "axios";
  
    const navigate = useNavigate();
    const location = useLocation();
  
    const user = {
      username: "",
      password: ""
    }
  
    async function postAPI() {
        await axios.post(
            `http://localhost:8000/login`,
            JSON.stringify(user)
            // {}, {
            //   auth: JSON.stringify(user)
            // }
        ).then(response => console.log(response));
        const from = ($location.state && $location.state.from) || "/";
        navigate(from, { replace: true });
    }

    function handleSubmit() {
      if (user.username.length === 0) {

      }
      if (user.password.length === 0) {

      }
      console.log(user);
      // Todo submit
      postAPI()
    }
  </script>
  
  <h3>Login</h3>
  <form on:submit|preventDefault={handleSubmit}>
    <input
      bind:value={user.username}
      type="text"
      name="username"
      placeholder="Username"
    />
    <br />
    <input
      bind:value={user.password}
      type="password"
      name="password"
      placeholder="Password"
    />
    <br />
    <button type="submit">Login</button>
  </form>