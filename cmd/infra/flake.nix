{
  inputs.app.url = github:defn/m/app-app-0.0.38?dir=app/app;
  outputs = inputs: inputs.app.goMain rec {
    src = ./.;

    extendBuild = ctx: {
      propagatedBuildInputs = [
        inputs.app.inputs.nodedev.defaultPackage.${ctx.system}
      ];
    };
  };
}
