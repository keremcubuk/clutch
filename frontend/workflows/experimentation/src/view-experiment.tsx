import React, { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import type { clutch as IClutch } from "@clutch-sh/api";
import { ExpansionPanel, ButtonGroup, client, Error, Row, Table, MetadataTable, TextField } from "@clutch-sh/core";
import { Container } from "@material-ui/core";
import styled from "styled-components";
// import { withRouter } from "react-router-dom";
import { Wizard } from "@clutch-sh/wizard";


interface ListExperimentsProps {
    columns: [string];
}


const Layout = styled(Container)`
  padding: 5% 0;
`;

export const Form = styled.form`
  align-items: center;
  display: flex;
  flex-direction: column;
  width: 100%;
`;

const StyledTextField = styled(TextField)`
  /* default */
  .MuiInput-underline:before {
    border-bottom: 0px solid green;
  }
  /* hover (double-ampersand needed for specificity reasons. */
  && .MuiInput-underline:hover:before {
    border-bottom: 0px solid lightblue;
  }
  /* focused */
  .MuiInput-underline:after {
    border-bottom: 0px solid red;
  }
`;

const ViewExperiment: React.FC<ListExperimentsProps> = ({ columns }) => {
    const [experiment, setExperiment] = useState("");
    const [error, setError] = useState("");

    const { id } = useParams();
    const navigate = useNavigate();

    enum Status {
        Scheduled = "scheduled",
        Running = "running",
        TerminatedManually = "terminated manually",
        Completed = "completed",
    };

    const experimentToStatus = function(experiment): Status {
        const now = new Date();
        if (experiment.startTime > now) {
            return Status.Scheduled;
        } else if (now > experiment.startTime && now < experiment.endTime) {
            return Status.Running;
        } else if (experiment.endTime < now && experiment.endTime < experiment.scheduledEndTime) {
            return Status.TerminatedManually;
        } else {
            return Status.Completed;
        }
    };

    function isCompleted(status: Status) {
        if (status == Status.TerminatedManually || status == Status.Completed) {
            return true;
        } else {
            return false;
        }
    };

    function capitalize(string) {
        return string.charAt(0).toUpperCase() + string.slice(1);
    };
    
    const dateToString = function(date: Date) {
        if (!date) {
            return "undefined"
        }

        return date.toISOString().substr(0,16)
    }

    const JSONToString = function(config) {
        return JSON.stringify(experiment, null, 4);
    };

    function buttons(experiment): ButtonProps[] {
        const goBack = function() { navigate("/experimentation/list") };

        const state = experimentToStatus(experiment);
        if (isCompleted(state)) {
            return [];
        }

        const title = state == Status.Running ? "Stop Experiment Run" : "Cancel Experiment Run";
        const destructiveButton =  {
            text: title,
            destructive: true,
            onClick: () => {
                console.log("call" + experiment.identifier);
                client.post("/v1/experiments/delete", {
                    ids: [experiment.identifier]
                })
                .then(response => {
                    goBack();
                })
                .catch(err => {
                    setError(err.response.statusText);
                });
            }
        };

        return [destructiveButton];
    }

    if (experiment === "") {
        client
          .post("/v1/experiments/get")
          .then(response => {
            let experiment = response?.data?.experiments[0];
            experiment.identifier = id;
            experiment.startTime = new Date(2020, 7, 20);
            experiment.endTime = new Date(2020, 12, 20);
            experiment.scheduledEndTime = new Date(2020, 12, 20);
            console.log(JSON.stringify(experiment, null, 2));
            setExperiment(experiment || []);
          })
          .catch(err => {
            setError(err.response.statusText);
          });

          return (
            <Form></Form>
          );
    }

    console.log(experiment);

    return (
        <Form>
            {error && <Error message={error} />}
            <TextField 
                label="Identifier"
                value={id}
                InputProps={{readOnly: true}}
            />
            <TextField 
                label="Status"
                value={capitalize(experimentToStatus(experiment))}
                InputProps={{readOnly: true}}
            />
            <TextField
                label="Start Time"
                defaultValue={dateToString(experiment.startTime)}
                type="datetime-local"
                InputProps={{readOnly: true}}
            />
            {isCompleted(experimentToStatus(experiment)) && 
            <TextField 
                label="End Time"
                defaultValue={dateToString(experiment.endTime)}
                type="datetime-local"
                InputProps={{readOnly: true}}
            />
            }
            <TextField 
                label="Scheduled End Time"
                defaultValue={dateToString(experiment.scheduledEndTime)}
                type="datetime-local"
                InputProps={{readOnly: true}}
            />
            <StyledTextField 
                multiline
                label="Config"
                defaultValue={JSONToString(experiment.config)}
                type="datetime-local"
                InputProps={{readOnly: true}}
            />
            <ButtonGroup
                buttons={buttons(experiment)}
            />
        </Form>
    )
}

export default ViewExperiment;
