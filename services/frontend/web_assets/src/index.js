import React, { useEffect, useState } from "react";
import ReactDOM from "react-dom";
import { ThemeProvider, CSSReset, Box, Heading, Text, Input, Button } from "@chakra-ui/core";
import { useDisclosure } from "@chakra-ui/core";
import { Stack } from "@chakra-ui/core";
import { Flex } from "@chakra-ui/core";
import { Textarea } from "@chakra-ui/core";
import { useToast } from "@chakra-ui/core";
import {
  Accordion,
  AccordionItem,
  AccordionHeader,
  AccordionPanel,
  AccordionIcon,
} from "@chakra-ui/core";
import {
  Drawer,
  DrawerBody,
  DrawerFooter,
  DrawerHeader,
  DrawerOverlay,
  DrawerContent,
  DrawerCloseButton,
} from "@chakra-ui/core";
import { FormControl } from "@chakra-ui/core";

const config = require('./appConfig.json');

const BACKEND_URL = config.BACKEND_URL;
const FEED_LABEL = config.FEED_LABEL;
const APP_TITLE = config.APP_TITLE;
const COMMENT_LABEL = config.COMMENT_LABEL;
const COMMENT_BUTTON_LABEL = config.COMMENT_BUTTON_LABEL;
const COMMENT_UNAVAILABLE_LABEL = config.COMMENT_UNAVAILABLE_LABEL;
const FEED_UNAVAILABLE_LABEL = config.FEED_UNAVAILABLE_LABEL;

function Feature({ title, desc, id, ...rest }) {
  const [error, setError] = useState(null);
  const [comments, setComments] = useState(null);

  useEffect(() => fetchComments(), [])

  function fetchComments() {
    fetch(BACKEND_URL + '/comments?event_id=' + id, {
      method: 'GET',
    })
    .then(res => res.json())
    .then(
      (result) => {
        setComments(result);
      },
      (error) => {
        setError(error);
      }
    )
  }

  function postComment(e) {
    e.preventDefault();
    const data = new FormData(e.target);
    fetch(BACKEND_URL + '/comments', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ 'commenter': 'supratik_das', 'content': data.get('content'), 'event_id': id })
    })
    .then(res => res.json())
    .then(data => {
      if (data.status == "success") {
        fetchComments();
      }
    })
  }

  return (
    <Box p={5} shadow="md" borderWidth="1px" {...rest}>
      <Heading fontSize="xl">{title}</Heading>
      <Text mt={4} mb={4}>{desc}</Text>
      <Accordion allowMultiple>
        <AccordionItem>
          <AccordionHeader _expanded={{ bg: "#1A365D", color: "white" }}>
            <Box flex="1" textAlign="left">
              {COMMENT_LABEL} ({comments === null ? 0 : comments.length})
            </Box>
            <AccordionIcon />
          </AccordionHeader>
          <AccordionPanel>
            <Stack spacing={3} align="center">
              { comments === null ? (<Text>{COMMENT_UNAVAILABLE_LABEL}</Text>) : comments.map(comment => (
                <Box p={2} w="100%" borderWidth="1px">
                  <Text>{comment.content}</Text>
                </Box>
              ))}
            </Stack>
            <form style={{width : '100%'}} onSubmit={postComment}>
              <FormControl mt={5}>
                <Input type="text" id="content" name="content" aria-describedby="text-helper-text" />
                <Button variantColor="teal" variant="solid" type="submit" mt={2}>  
                  {COMMENT_BUTTON_LABEL}
                </Button>
              </FormControl>
            </form>
          </AccordionPanel>
        </AccordionItem>
      </Accordion>
    </Box>
  );
}

function NewEventDrawer(props) {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const toast = useToast();
  const btnRef = React.useRef();

  function postEvent(e) {
    e.preventDefault();
    const data = new FormData(e.target);
    fetch(BACKEND_URL + '/events', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ 'name': data.get('name'), 'description': data.get('description') })
    })
    .then(res => res.json())
    .then(data => {
      if (data.status === "success") {
        props.callBack(true);
        toast({
          title: "Event created.",
          description: "",
          status: "success",
          duration: 5000,
          isClosable: true,
        })
      } else {
        toast({
          title: "Event creation failed",
          description: "",
          status: "error",
          duration: 5000,
          isClosable: true,
        })
      }
    })
  }

  return (
    <>
      <Button ref={btnRef} variantColor="teal" onClick={onOpen}>
        Create an {FEED_LABEL}
      </Button>
      <Drawer
        isOpen={isOpen}
        placement="top"
        onClose={onClose}
        finalFocusRef={btnRef}
      >
        <DrawerOverlay />
        <DrawerContent>
          <form style={{width : '100%'}} onSubmit={postEvent}>
            <DrawerCloseButton />
            <DrawerHeader>Create an {FEED_LABEL}</DrawerHeader>

            <DrawerBody>
              <Input name="name" placeholder="Event Name..." />
              <Textarea name="description" mt={5} placeholder="Event Description..." />
            </DrawerBody>

            <DrawerFooter>
              <Button variant="outline" mr={3} onClick={onClose}>
                Cancel
              </Button>
              <Button type="submit" color="blue">Save</Button>
            </DrawerFooter>
          </form>
        </DrawerContent>
      </Drawer>
    </>
  );
}

function App() {
  const [error, setError] = useState(null);
  const [events, setEvents] = useState(null);
  const [reload, setReload] = useState(false);

  useEffect(() => 
    fetch(BACKEND_URL + '/events', {
      method: 'GET',
    })
    .then(res => res.json())
    .then(
      (result) => {
        setEvents(result);
      },
      (error) => {
        setError(error);
      }
  ), [])

  return (
    <ThemeProvider>
      <CSSReset />
      <Box bg="#1A365D" w="100%" p={4} color="white">
        <Flex align="center" justify="space-between">
          <Heading as="h2" size="xl">
            {APP_TITLE}
          </Heading>
          <NewEventDrawer callBack={setReload}></NewEventDrawer>
        </Flex>
      </Box>
      <Stack pl={20} pr={20} pt={10} spacing={8}>
        { events === null ? (<Text>Loading</Text>) : events.map(event => (
            <Feature
            title={event.name}
            desc={event.description}
            id={event.id}
          />
          ))
         }
    </Stack>
    </ThemeProvider>
  );
}

const rootElement = document.getElementById("root");
ReactDOM.render(<App />, rootElement);