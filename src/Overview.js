import * as React from "react";
import InitialStateCard from "./InitialStateCard";
import UpdateCard from "./UpdateCard";
import { Classes, H3, Icon, Intent, InputGroup, Tab, Tabs, Card } from "@blueprintjs/core";

class Overview extends React.PureComponent {
  render() {
    const {selectTabID, selectedTabID} = this.props; 
    return (
        <Tabs vertical animate class="Tabs" id="TabsExample">
          <Tab
            id="config"
            title={
              <div>
                <Icon
                  className="TabIcon"
                  iconSize={15}
                  icon="wrench"
                  intent={Intent.PRIMARY}
                />
                Config
              </div>
            }
            panel={<ConfigPanel />}
            panelClassName="ember-panel"
          />
            
          <Tab id="initial" title={
              <div>
                <Icon
                  className="TabIcon"
                  iconSize={15}
                  icon="diagram-tree"
                  intent={Intent.PRIMARY}
                />
                Initial State
              </div>
            }  panel={<InitialStateCard />} />
          <Tab id="update" title={
              <div>
                <Icon
                  className="TabIcon"
                  iconSize={15}
                  icon="diagram-tree"
                  intent={Intent.PRIMARY}
                />
                Latest Update
              </div>
            }  panel={<InitialStatePanel />} />
          <Tabs.Expander />
        </Tabs>
    );
  }
}

const InitialStatePanel = () => (
  <div>
    <H3>Example panelReact</H3>
    <p className={Classes.RUNNING_TEXT}>
      Lots of people use React as the V in MVC. Since React makes no assumptions
      about the rest of your technology stack, it's easy to try it out on a
      small feature in an existing project.
    </p>
  </div>
);

const LatestUpdatePanel = () => (
  <div>
    <H3>Example panelAngular</H3>
    <p className={Classes.RUNNING_TEXT}>
      HTML is great for declaring static documents, but it falters when we try
      to use it for declaring dynamic views in web-applications. AngularJS lets
      you extend HTML vocabulary for your application. The resulting environment
      is extraordinarily expressive, readable, and quick to develop.
    </p>
  </div>
);

const ConfigPanel = () => (
  <div>
    <H3>Example panelEmber</H3>
    <p className={Classes.RUNNING_TEXT}>
      Ember.js is an open-source JavaScript application framework, based on the
      model-view-controller (MVC) pattern. It allows developers to create
      scalable single-page web applications by incorporating common idioms and
      best practices into the framework. What is your favorite JS framework?
    </p>
    <input className={Classes.INPUT} type="text" />
  </div>
);

export default Overview;
