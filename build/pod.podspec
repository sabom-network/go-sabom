Pod::Spec.new do |spec|
  spec.name         = 'Gsab'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/sabom-network/go-sabom'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Sabom Client'
  spec.source       = { :git => 'https://github.com/sabom-network/go-sabom.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Gsab.framework'

	spec.prepare_command = <<-CMD
    curl https://gsabstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Gsab.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
