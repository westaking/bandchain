module Styles = {
  open Css;

  let container =
    style([
      backgroundColor(Colors.white),
      boxShadow(Shadow.box(~x=`zero, ~y=`px(2), ~blur=`px(10), Css.rgba(0, 0, 0, 0.08))),
    ]);
  let header =
    style([
      backgroundColor(Colors.white),
      padding2(~v=`zero, ~h=`px(20)),
      borderBottom(`px(1), `solid, Colors.gray4),
      boxShadow(Shadow.box(~x=`zero, ~y=`px(2), ~blur=`px(10), Css.rgba(0, 0, 0, 0.08))),
    ]);

  let buttonContainer = active =>
    style([
      height(`px(40)),
      display(`inlineFlex),
      justifyContent(`center),
      alignItems(`center),
      cursor(`pointer),
      padding2(~v=Spacing.md, ~h=`px(20)),
      borderBottom(`pxFloat(1.5), `solid, active ? Colors.bandBlue : Colors.white),
      textShadow(Shadow.text(~blur=`pxFloat(active ? 1. : 0.), Colors.bandBlue)),
    ]);

  let childrenContainer = style([backgroundColor(Colors.blueGray1)]);
};

let button = (~name, ~route, ~active) => {
  <Link key=name className={Styles.buttonContainer(active)} route>
    <Text
      value=name
      weight=Text.Regular
      size=Text.Md
      color={active ? Colors.bandBlue : Colors.gray6}
    />
  </Link>;
};

type t = {
  name: string,
  route: Route.t,
};

[@react.component]
let make = (~tabs: array(t), ~currentRoute, ~children) => {
  <div className=Styles.container>
    <div className=Styles.header>
      <Row>
        {tabs
         ->Belt.Array.map(({name, route}) =>
             button(~name, ~route, ~active=route == currentRoute)
           )
         ->React.array}
      </Row>
    </div>
    <div className=Styles.childrenContainer> children </div>
  </div>;
};
