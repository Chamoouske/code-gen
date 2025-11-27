package dependencyinjection

import (
	"code-gen/internal/domain"
	"code-gen/internal/infra/converter"
)

func addConverters(sl domain.SupportedLangages) {
	sl.AddSupport("java", converter.NewJavaConverter())
}
