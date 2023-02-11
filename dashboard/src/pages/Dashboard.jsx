import * as React from "react";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Link from "@mui/material/Link";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import { Link as RouterLink } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { styled } from "@mui/material/styles";
import { Divider } from "@mui/material";

const Row = styled("div")({
  display: "flex",
  flexDirection: "row",
  justifyContent: "space-between",
  width: "100%",
  alignItems: "center",
});

const ServicesLayout = styled("div")({
  display: "flex",
  flexDirection: "column",
  justifyContent: "space-around",
  alignItems: "center",
  width: "100%",
  margin: "1rem",
  gap: "1.6rem",
});

const Service = (props) => {
  return (
    <Box
      sx={{
        // border: '1px solid #E9F8F9',
        backgroundColor: "#C0EEF2",
        color: "#181823",
        display: "flex",
        flexDirection: "column",
        alignItems: "flex-start",
        justifyContent: "center",
        height: "100%",
        width: "100%",
        p: 2,
        borderRadius: "4px",
        transition: "all 0.1s ease-in-out",
        "&:hover": {
          backgroundColor: "#537FE7",
          boxShadow: "0 0 16px 8px #C0EEF222",
          color: "#fff",
        },
        cursor: "pointer",
      }}
    >
      <Row>
        <Typography component="h5" variant="h6">
          {props.service.name}
        </Typography>
        <Typography
          component="h5"
          variant="h6"
          sx={{
            fontSize: "0.8rem",
          }}
        >
          {props.service.dateCreated}
        </Typography>
      </Row>

      <Divider sx={{ width: "100%", my: 1 }} color="#181823" />
      <Typography component="p" variant="body2">
        {props.service.description}
      </Typography>
      <Link
        component={RouterLink}
        to={`/app/${props.service.serviceId}/v`}
        variant="body2"
        style={{
          textDecoration: "none",
          backgroundColor: "#2ecc71",
          padding: "0.5rem",
          borderRadius: "0.2rem",
          alignSelf: "flex-end",
          color: "#181823",
        }}
      >
        {"View Service"}
      </Link>
    </Box>
  );
};

const ServiceCreator = (props) => {
  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    login({
      email: data.get("email"),
      password: data.get("password"),
    });
  };
  return (
    <Box
      component="form"
      onSubmit={handleSubmit}
      noValidate
      sx={{ border: "1px solid #C0EEF2", padding: "0 1rem 1rem", borderRadius: "4px", animation: "slideIn 1s ease-in-out", display: "flex", flexDirection: "column", width: "100%", cursor: "pointer"}}
      
    >
    <Row sx={{
      alignItems: "center",
      marginTop: "1rem",
    }}>
    <Typography component="h1" variant="h5"  >
      Create a new service
    </Typography>
      <Typography component="h1" variant="h5" sx={{
        fontSize: "1.2rem",
        backgroundColor: "#e74c3c",
        width: "fit-content",
        padding: "0.2rem 0.8rem",
        alignSelf: "flex-end",
      }} onClick={() => {
        props.closeCreator();
      }}>
      X
      </Typography>
      </Row>
      <TextField
        margin="normal"
        required
        fullWidth
        id="email"
        label="Service Name"
        name="email"
        // autoComplete="email"
        autoFocus
      />
      <TextField
        margin="normal"
        // required
        fullWidth
        id="email"
        label="Service Description"
        name="desc"
        // autoComplete="email"
        // autoFocus
      />
      {/* <TextField
        margin="normal"
        required
        fullWidth
        name="password"
        label="Password"
        type="password"
        id="password"
        autoComplete="current-password"
      /> */}
      <Button
        type="submit"
        fullWidth
        variant="contained"
        sx={{ mt: 3, mb: 2 }}
        color="success"
      >
        Create new Service
      </Button>
      {/* <Grid container>
        <Grid item>
          <RouterLink to="/signup">
            <Link href="#" variant="body2">
              {"New here? Sign Up"}
            </Link>
          </RouterLink>
        </Grid>
      </Grid> */}
    </Box>
  );
};

const Services = (props) => {
  const [creating, setCreating] = React.useState(false);

  return (
    <ServicesLayout>
      {!creating ? (
        <Typography
          component="p"
          variant="body2"
          sx={{
            fontSize: "2rem",
            cursor: "pointer",
            border: "1px solid #C0EEF2",
            padding: "0 1rem",
            borderRadius: "4px",
          }}
          onClick={() => setCreating(true)}
        >
          {"+"}
        </Typography>
      ) : (
        <ServiceCreator closeCreator={() => {
          setCreating(false);
        }}></ServiceCreator>
      )}
      {props.children}
    </ServicesLayout>
  );
};

export const Dashboard = () => {
  const { user } = useAuth();

  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    login({
      email: data.get("email"),
      password: data.get("password"),
    });
  };
  const services = [
    {
      name: "Service 1",
      description: "This is service 1",
      serviceId: 1,
      dateCreated: "2023-02-11",
    },
    {
      name: "Service 2",
      description: "This is service 2",
      serviceId: 2,
      dateCreated: "2023-02-11",
    },
    {
      name: "Service 3",
      description: "This is service 3",
      serviceId: 3,
      dateCreated: "2023-02-11",
    },
    {
      name: "Service 4",
      description: "This is service 4",
      serviceId: 4,
      dateCreated: "2023-02-11",
    },
  ];
  return (
    <Container
      component="main"
      maxWidth="sm"
      sx={{
        marginTop: "1vh",
        display: "flex",
        flexDirection: "column",
        justifyContent: "flex-start",
        height: "88vh",
        alignItems: "center",
        overflowY: "scroll",
        scrollbarWidth: 0,
        "&::-webkit-scrollbar": {
          display: "none",
        },
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          width: "100%",
        }}
      >
        <Typography
          component="h1"
          variant="h4"
          sx={{
            fontWeight: "bold",
            margin: "1rem 0",
          }}
        >
          Your services
        </Typography>
        <Services>
          {services.map((service) => (
            <Service service={service} />
          ))}
        </Services>
      </Box>
    </Container>
  );
};
